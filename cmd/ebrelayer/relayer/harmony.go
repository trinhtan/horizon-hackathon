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
	logs := make(chan ctypes.Log)

	// Start BridgeBank subscription, prepare contract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startHarmonyContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := contract.LoadHarmonyABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.HmyLogLock.String()].Id().Hex()

	// Start ethereumBridge subscription, prepare contract ABI and HmyLogNewUnlockClaim event signature
	_, subEthereumBridge := sub.startHarmonyContractEventSub(logs, client, txs.EthereumBridge)
	ethereumBridgeContractABI := contract.LoadHarmonyABI(txs.EthereumBridge)
	_ = ethereumBridgeContractABI.Events[types.HmyLogNewUnlockClaim.String()].Id().Hex()

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
				// case eventLogNewUnlockClaimSignature:
				// 	err = sub.handleHarmonyLogNewUnlockClaim(ethereumBridgeAddress, ethereumBridgeContractABI,
				// 		types.HmyLogNewUnlockClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
				continue
			}
		}
	}

	// query := ethereum.FilterQuery{
	// 	Addresses: []common.Address{sub.BridgeRegistry},
	// }

	// contractSub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	select {
	// 	case err := <-contractSub.Err():
	// 		log.Fatal(err)
	// 	case vLog := <-logs:

	// 		fmt.Println("Data: ", vLog.Data) // pointer to event log
	// 		event := types.HarmonyLogBridgeBankSetEvent{}
	// 		err := bridgeRegistryContractABI.Unpack(&event, types.LogBridgeBankSet.String(), vLog.Data)
	// 		if err != nil {
	// 			sub.Logger.Error("error unpacking: %v", err)
	// 		}
	// 		sub.Logger.Info(event.String())
	// 	}
	// }
}

// startContractEventSub : starts an event subscription on the specified Ethereum contract
func (sub HarmonySub) startHarmonyContractEventSub(logs chan ctypes.Log, client *hmyclient.Client,
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
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via contract ABI
	event := types.HmyLogLockEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = bridgeBankAddress
	event.HarmonyChainID = clientChainID

	sub.Logger.Info(event.String())

	// return nil

	// Add the event to the record
	types.NewHarmonyEventWrite(cLog.TxHash.Hex(), event)

	unlockClaim, err := txs.HarmonyEventToEthereumBridgeClaim(&event)
	if err != nil {
		return err
	}

	return txs.RelayUnlockClaimToEthereum(sub.EthereumProvider, sub.EthereumBridgeRegistry, types.HmyLogLock, unlockClaim, sub.PrivateKey)
}

// // Unpacks a handleHarmonyLogNewUnlockClaim event, builds a new OracleClaim, and relays it to Ethereum
// func (sub HarmonySub) handleHarmonyLogNewUnlockClaim(contractAddress common.Address, contractABI abi.ABI,
// 	eventName string, cLog ctypes.Log) error {
// 	// Parse the event's attributes via contract ABI
// 	event := types.HmyLogNewUnlockClaimEvent{}
// 	err := contractABI.Unpack(&event, eventName, cLog.Data)
// 	if err != nil {
// 		sub.Logger.Error("error unpacking: %v", err)
// 	}
// 	sub.Logger.Info(event.String())

// 	// oracleClaim, err := txs.UnlockClaimToSignedOracleClaim(event, sub.PrivateKey)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// return txs.RelayOracleClaimToEthereum(sub.HarmonyProvider, contractAddress, types.EthLogNewUnlockClaim,
// 	// 	oracleClaim, sub.PrivateKey)
// 	return nil
// }

// func main() {
// 	// client, err := hmyclient.Dial("wss://rinkeby.infura.io/ws/v3/469a0721f318401584ab3639fb548d63")
// 	url := "wss://ws.s0.b.hmny.io"
// 	client, err := hmyclient.Dial(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("url: ", url, "\n")

// 	fmt.Println("client: ", client, "\n")

// 	clientChainID, err := client.NetworkID(context.Background())

// 	fmt.Println("ChainID: ", clientChainID, "\n")

// 	contractAddress := common.HexToAddress("0x261d28bfb27F4B69d5a72bEE1b8C1F1C30D772F8")

// 	fmt.Println("contract address: ", contractAddress, "\n")

// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{contractAddress},
// 	}

// 	fmt.Println("query: ", query, "\n")

// 	logs := make(chan types.Log)

// 	fmt.Println("logs: ", logs, "\n")

// 	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:

// 			fmt.Println("Data: ", vLog.Data) // pointer to event log

// 		}
// 	}
// }
