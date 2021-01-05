package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
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
	HmyProvider    string
	BridgeRegistry common.Address
	ValidatorName  string
	PrivateKey     *ecdsa.PrivateKey
	Logger         tmLog.Logger
}

// NewHarmonySub initializes a new HarmonySub
func NewHarmonySub(inBuf io.Reader, validatorMoniker,
	hmyProvider string, bridgeRegistry common.Address, privateKey *ecdsa.PrivateKey,
	logger tmLog.Logger) (HarmonySub, error) {

	return HarmonySub{
		HmyProvider:    hmyProvider,
		BridgeRegistry: bridgeRegistry,
		ValidatorName:  "validator",
		PrivateKey:     privateKey,
		Logger:         logger,
	}, nil
}

// Start an Ethereum chain subscription
func (sub HarmonySub) Start() {
	client, err := SetupWebsocketHmyClient(sub.HmyProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Harmony websocket with provider:", sub.HmyProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println("Harmony ChainID: ", clientChainID)

	// We will check logs for new events
	logs := make(chan ctypes.Log)

	// Start BridgeBank subscription, prepare contract ABI and LockLog event signature
	_, subBridgeBank := sub.startHmyContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := contract.LoadABI(txs.BridgeBank)
	_ = bridgeBankContractABI.Events[types.HmyLogLock.String()].Id().Hex()

	// Start harmonyBridge subscription, prepare contract ABI and LogNewProphecyClaim event signature
	_, subHarmonyBridge := sub.startHmyContractEventSub(logs, client, txs.HarmonyBridge)
	harmonyBridgeContractABI := contract.LoadABI(txs.HarmonyBridge)
	_ = harmonyBridgeContractABI.Events[types.HmyLogNewUnlockClaim.String()].Id().Hex()

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
			// var err error
			// switch vLog.Topics[0].Hex() {
			// case eventLogLockSignature:
			// 	err = sub.handleEthereumLockEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
			// 		types.LogLock.String(), vLog)
			// case eventLogNewUnlockClaimSignature:
			// 	err = sub.handleLogNewUnlockClaim(harmonyBridgeAddress, harmonyBridgeContractABI,
			// 		types.LogNewUnlockClaim.String(), vLog)
			// }
			// // TODO: Check local events store for status, if retryable, attempt relay again
			// if err != nil {
			// 	sub.Logger.Error(err.Error())
			// }
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
func (sub HarmonySub) startHmyContractEventSub(logs chan ctypes.Log, client *hmyclient.Client,
	contractName txs.ContractRegistry) (common.Address, ethereum.Subscription) {
	// Get the contract address for this subscription
	subContractAddress, err := txs.HmyGetAddressFromBridgeRegistry(client, sub.BridgeRegistry, contractName)
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
