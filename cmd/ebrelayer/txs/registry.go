package txs

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	bridgeregistry "github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/contract/generated/bindings/bridgeregistry"
	"github.com/trinhtan/horizon-hackathon/hmyclient"
)

// TODO: Update BridgeRegistry contract so that all bridge contract addresses can be queried
//		in one transaction. Then refactor ContractRegistry to a map and store it under new
//		Relayer struct.

// ContractRegistry is an enum for the bridge contract types
type ContractRegistry byte

const (
	// Valset valset contract
	Valset ContractRegistry = iota + 1
	// Oracle oracle contract
	Oracle
	// BridgeBank bridgeBank contract
	BridgeBank
	// HarmonyBridge cosmosBridge contract
	HarmonyBridge
	// BridgeRegistry contract
	BridgeRegistry
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "harmonybridge", "bridgeregistry"}[d-1]
}

// EthGetAddressFromBridgeRegistry queries the requested contract address from the BridgeRegistry contract
func EthGetAddressFromBridgeRegistry(client *ethclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	sender, err := LoadSender()
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up CallOpts auth
	auth := bind.CallOpts{
		Pending:     true,
		From:        sender,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}

	// Initialize BridgeRegistry instance
	registryInstance, err := bridgeregistry.NewBridgeRegistry(registry, client)
	if err != nil {
		log.Fatal(err)
	}

	var address common.Address
	switch target {
	case Valset:
		address, err = registryInstance.Valset(&auth)
	case Oracle:
		address, err = registryInstance.Oracle(&auth)
	case BridgeBank:
		address, err = registryInstance.BridgeBank(&auth)
	case HarmonyBridge:
		address, err = registryInstance.HarmonyBridge(&auth)
	default:
		panic("invalid target contract address")
	}

	if err != nil {
		log.Fatal(err)
	}

	return address, nil
}

// HmyGetAddressFromBridgeRegistry queries the requested contract address from the BridgeRegistry contract
func HmyGetAddressFromBridgeRegistry(client *hmyclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	// sender, err := LoadSender()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Set up CallOpts auth
	// auth := bind.CallOpts{
	// 	Pending:     true,
	// 	From:        sender,
	// 	BlockNumber: header.Number,
	// 	Context:     context.Background(),
	// }

	// // Initialize BridgeRegistry instance
	// registryInstance, err := bridgeregistry.NewBridgeRegistry(registry, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var address common.Address
	switch target {
	case Valset:
		// address, err = registryInstance.Valset(&auth)
		address = common.HexToAddress("0x3EBB943c851040C2C56b2020210e8472a547ee1D")
	case Oracle:
		// address, err = registryInstance.Oracle(&auth)
		address = common.HexToAddress("0x3EBB943c851040C2C56b2020210e8472a547ee1D")
	case BridgeBank:
		// address, err = registryInstance.BridgeBank(&auth)
		address = common.HexToAddress("0x3EBB943c851040C2C56b2020210e8472a547ee1D")
	case HarmonyBridge:
		// address, err = registryInstance.HarmonyBridge(&auth)
		address = common.HexToAddress("0x3EBB943c851040C2C56b2020210e8472a547ee1D")
	default:
		panic("invalid target contract address")
	}

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return address, nil
}
