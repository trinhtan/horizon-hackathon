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

	hmytypes "github.com/harmony-one/harmony/core/types"
	tmLog "github.com/tendermint/tendermint/libs/log"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/contract"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/txs"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/types"
	"github.com/trinhtan/horizon-hackathon/hmyclient"
)

// HarmonySub is an Ethereum listener that can relay txs to Cosmos and Ethereum
type HarmonySub struct {
	HarmonyProvider        string
	EthereumProvider       string
	HarmonyBridgeRegistry  common.Address
	EthereumBridgeRegistry common.Address
	ValidatorName          string
	PrivateKey             *ecdsa.PrivateKey
	Logger                 tmLog.Logger
}

// NewHarmonySub initializes a new HarmonySub
func NewHarmonySub(inBuf io.Reader, validatorMoniker,
	harmonyProvider string, ethereumProvider string, harmonyBridgeRegistry common.Address, ethereumBridgeRegistry common.Address, privateKey *ecdsa.PrivateKey,
	logger tmLog.Logger) (HarmonySub, error) {

	return HarmonySub{
		HarmonyProvider:        harmonyProvider,
		EthereumProvider:       ethereumProvider,
		HarmonyBridgeRegistry:  harmonyBridgeRegistry,
		EthereumBridgeRegistry: ethereumBridgeRegistry,
		ValidatorName:          "validator",
		PrivateKey:             privateKey,
		Logger:                 logger,
	}, nil
}

// Start an Ethereum chain subscription
func (sub HarmonySub) Start() {
	client, err := SetupWebsocketHmyClient(sub.HarmonyProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Harmony websocket with provider:", sub.HarmonyProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println("Harmony ChainID: ", clientChainID)

	// We will check logs for new events
	logs := make(chan hmytypes.Log)

	// Start BridgeBank subscription, prepare contract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startHarmonyContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := contract.LoadHarmonyABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.HmyLogLock.String()].ID.Hex()

	// Start ethereumBridge subscription, prepare contract ABI and HmyLogNewUnlockClaim event signature
	_, subEthereumBridge := sub.startHarmonyContractEventSub(logs, client, txs.EthereumBridge)
	ethereumBridgeContractABI := contract.LoadHarmonyABI(txs.EthereumBridge)
	eventLogNewUnlockClaimSignature := ethereumBridgeContractABI.Events[types.HmyLogNewUnlockClaim.String()].ID.Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeBank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subEthereumBridge.Err():
			sub.Logger.Error(err.Error())
		// vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventLogLockSignature:
				err = sub.handleHarmonyLogLockEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI, types.HmyLogLock.String(), vLog)
			case eventLogNewUnlockClaimSignature:
				err = sub.handleHarmonyLogNewUnlockClaim(sub.HarmonyBridgeRegistry, ethereumBridgeContractABI,
					types.HmyLogNewUnlockClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
				continue
			}
		}
	}

}

// startContractEventSub : starts an event subscription on the specified Ethereum contract
func (sub HarmonySub) startHarmonyContractEventSub(logs chan hmytypes.Log, client *hmyclient.Client,
	contractName txs.ContractRegistry) (common.Address, ethereum.Subscription) {
	// Get the contract address for this subscription
	subContractAddress, err := txs.HmyGetAddressFromBridgeRegistry(client, sub.HarmonyBridgeRegistry, contractName)
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
	sub.Logger.Info(fmt.Sprintf("Harmony side - Subscribed to %v contract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleHarmonyLogLockEvent unpacks an Ethereum lock token event, and relays a tx to Harmony
func (sub HarmonySub) handleHarmonyLogLockEvent(clientChainID *big.Int, bridgeBankAddress common.Address,
	contractABI abi.ABI, eventName string, cLog hmytypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.HmyLogLockEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = bridgeBankAddress
	event.HarmonyChainID = clientChainID

	sub.Logger.Info(event.String())

	unlockClaim, err := txs.HarmonyEventToEthereumBridgeClaim(&event)
	if err != nil {
		return err
	}

	return txs.RelayUnlockClaimToEthereum(sub.EthereumProvider, sub.EthereumBridgeRegistry, types.HmyLogLock, unlockClaim, sub.PrivateKey)
}

// Unpacks a HmyLogNewUnlockClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub HarmonySub) handleHarmonyLogNewUnlockClaim(contractAddress common.Address, contractABI abi.ABI,
	eventName string, hLog hmytypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.HmyLogNewUnlockClaimEvent{}
	err := contractABI.Unpack(&event, eventName, hLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.HmyUnlockClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToHarmony(sub.HarmonyProvider, contractAddress, types.HmyLogNewUnlockClaim,
		oracleClaim, sub.PrivateKey)
}
