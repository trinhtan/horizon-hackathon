package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
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
	Cdc                     *codec.Codec
	EthProvider             string
	RegistryContractAddress common.Address
	ValidatorName           string
	PrivateKey              *ecdsa.PrivateKey
	Logger                  tmLog.Logger
}

// NewEthereumSub initializes a new EthereumSub
func NewEthereumSub(inBuf io.Reader, rpcURL string, cdc *codec.Codec, validatorMoniker, chainID,
	ethProvider string, registryContractAddress common.Address, privateKey *ecdsa.PrivateKey,
	logger tmLog.Logger) (EthereumSub, error) {

	return EthereumSub{
		Cdc:                     cdc,
		EthProvider:             ethProvider,
		RegistryContractAddress: registryContractAddress,
		ValidatorName:           "validator",
		PrivateKey:              privateKey,
		Logger:                  logger,
	}, nil
}

// Start an Ethereum chain subscription
func (sub EthereumSub) Start() {
	client, err := SetupWebsocketEthClient(sub.EthProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Ethereum websocket with provider:", sub.EthProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}

	// We will check logs for new events
	logs := make(chan ctypes.Log)

	// Start BridgeBank subscription, prepare contract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := contract.LoadABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.LogLock.String()].Id().Hex()

	// Start harmonyBridge subscription, prepare contract ABI and LogNewProphecyClaim event signature
	harmonyBridgeAddress, subHarmonyBridge := sub.startContractEventSub(logs, client, txs.HarmonyBridge)
	harmonyBridgeContractABI := contract.LoadABI(txs.HarmonyBridge)
	eventLogNewUnlockClaimSignature := harmonyBridgeContractABI.Events[types.LogNewUnlockClaim.String()].Id().Hex()

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
				err = sub.handleEthereumLockEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogLock.String(), vLog)
			case eventLogNewUnlockClaimSignature:
				err = sub.handleLogNewUnlockClaim(harmonyBridgeAddress, harmonyBridgeContractABI,
					types.LogNewUnlockClaim.String(), vLog)
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
	subContractAddress, err := txs.GetAddressFromBridgeRegistry(client, sub.RegistryContractAddress, contractName)
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
	sub.Logger.Info(fmt.Sprintf("Subscribed to %v contract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleEthereumLockEvent unpacks an Ethereum lock token event, and relays a tx to Harmony
func (sub EthereumSub) handleEthereumLockEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.EthereumLogLockEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = contractAddress
	event.EthereumChainID = clientChainID

	// if eventName == types.LogBurn.String() {
	// 	event.ClaimType = ethbridge.BurnText
	// } else {
	// 	event.ClaimType = ethbridge.LockText
	// }
	sub.Logger.Info(event.String())

	// Add the event to the record
	types.NewEventWrite(cLog.TxHash.Hex(), event)

	// prophecyClaim, err := txs.EthereumEventToEthBridgeClaim(sub.ValidatorAddress, &event)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(prophecyClaim)

	return nil
	// return txs.RelayToCosmos(sub.Cdc, sub.ValidatorName, &prophecyClaim, sub.CliCtx, sub.TxBldr)
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EthereumSub) handleLogNewUnlockClaim(contractAddress common.Address, contractABI abi.ABI,
	eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.EthereumLogNewUnlockClaimEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.ProphecyClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToEthereum(sub.EthProvider, contractAddress, types.LogNewUnlockClaim,
		oracleClaim, sub.PrivateKey)
	// return nil
}
