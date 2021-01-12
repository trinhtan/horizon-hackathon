package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"math/big"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/contract"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/txs"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/types"
)

// TODO: Move relay functionality out of EthereumSub into a new Relayer parent struct

// EthereumSub is an Ethereum listener that can relay txs to Cosmos and Ethereum
type EthereumSub struct {
	EthereumProvider       string
	HarmonyProvider        string
	EthereumBridgeRegistry common.Address
	HarmonyBridgeRegistry  common.Address
	ValidatorName          string
	PrivateKey             *ecdsa.PrivateKey
	Logger                 tmLog.Logger
}

// NewEthereumSub initializes a new EthereumSub
func NewEthereumSub(inBuf io.Reader, validatorMoniker, ethereumProvider string, harmonyProvider string,
	ethereumBridgeRegistry common.Address, harmonyBridgeRegistry common.Address, privateKey *ecdsa.PrivateKey, logger tmLog.Logger) (EthereumSub, error) {
	return EthereumSub{
		EthereumProvider:       ethereumProvider,
		HarmonyProvider:        harmonyProvider,
		EthereumBridgeRegistry: ethereumBridgeRegistry,
		HarmonyBridgeRegistry:  harmonyBridgeRegistry,
		ValidatorName:          "validator",
		PrivateKey:             privateKey,
		Logger:                 logger,
	}, nil
}

// Start an Ethereum chain subscription
func (sub EthereumSub) Start() {
	client, err := SetupWebsocketEthClient(sub.EthereumProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Ethereum websocket with provider:", sub.EthereumProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println("Ethereum ChainID: ", clientChainID)

	// We will check logs for new events
	logs := make(chan ctypes.Log)

	// Start BridgeBank subscription, prepare contract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := contract.LoadEthereumABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.EthLogLock.String()].ID.Hex()

	// Start harmonyBridge subscription, prepare contract ABI and EthLogNewUnlockClaim event signature
	_, subHarmonyBridge := sub.startContractEventSub(logs, client, txs.HarmonyBridge)
	harmonyBridgeContractABI := contract.LoadEthereumABI(txs.HarmonyBridge)
	eventLogNewUnlockClaimSignature := harmonyBridgeContractABI.Events[types.EthLogNewUnlockClaim.String()].ID.Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeBank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subHarmonyBridge.Err():
			sub.Logger.Error(err.Error())
		// vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventLogLockSignature:
				err = sub.handleEthereumLogLockEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.EthLogLock.String(), vLog)
			case eventLogNewUnlockClaimSignature:
				err = sub.handleEthLogNewUnlockClaim(sub.EthereumBridgeRegistry, harmonyBridgeContractABI,
					types.EthLogNewUnlockClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
}

// startContractEventSub : starts an event subscription on the specified Ethereum contract
func (sub EthereumSub) startContractEventSub(logs chan ctypes.Log, client *ethclient.Client,
	contractName txs.ContractRegistry) (common.Address, ethereum.Subscription) {
	// Get the contract address for this subscription
	subContractAddress, err := txs.EthGetAddressFromBridgeRegistry(client, sub.EthereumBridgeRegistry, contractName)
	if err != nil {
		sub.Logger.Error(err.Error())
	}

	// We need the address in []bytes for the query
	subQuery := ethereum.FilterQuery{
		Addresses: []common.Address{subContractAddress},
	}

	// Start the contract subscription
	contractSub, err := client.SubscribeFilterLogs(context.Background(), subQuery, logs)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	sub.Logger.Info(fmt.Sprintf("Ethereum side - Subscribed to %v contract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleEthereumLogLockEvent unpacks an Ethereum lock token event, and relays a tx to Harmony
func (sub EthereumSub) handleEthereumLogLockEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.EthLogLockEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = contractAddress
	event.EthereumChainID = clientChainID
	sub.Logger.Info(event.String())

	// Add the event to the record
	types.NewEventWrite(cLog.TxHash.Hex(), event)

	unlockClaim, err := txs.EthereumEventToHarmonyBridgeClaim(&event)
	if err != nil {
		return err
	}

	return txs.RelayUnlockClaimToHarmony(sub.HarmonyProvider, sub.HarmonyBridgeRegistry, types.EthLogLock, unlockClaim, sub.PrivateKey)
}

// Unpacks a EthLogNewUnlockClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EthereumSub) handleEthLogNewUnlockClaim(contractAddress common.Address, contractABI abi.ABI,
	eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.EthLogNewUnlockClaimEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.EthUnlockClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToEthereum(sub.EthereumProvider, contractAddress, types.EthLogNewUnlockClaim,
		oracleClaim, sub.PrivateKey)
}
