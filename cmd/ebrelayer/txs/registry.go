package txs

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	ethereumbridgeregistry "github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/contract/generated/ethereum/bindings/bridgeregistry"
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
	// HarmonyBridge  contract
	HarmonyBridge
	// EthereumBridge  contract
	EthereumBridge
	// BridgeRegistry contract
	BridgeRegistry
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "harmonybridge", "ethereumbridge", "bridgeregistry"}[d-1]
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
	registryInstance, err := ethereumbridgeregistry.NewBridgeRegistry(registry, client)
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
		address = common.HexToAddress("0x294cBe28EE7A28c0f9B0473d527dfEd103a5a286")
	case Oracle:
		// address, err = registryInstance.Oracle(&auth)
		address = common.HexToAddress("0xA0386fbd75E368001B8eF0b3F33e12261b6c787F")
	case EthereumBridge:
		// address, err = registryInstance.HarmonyBridge(&auth)
		address = common.HexToAddress("0xb5b7C741235A8eebbAc0Fe5782a0A2c235fA12dA")
	case BridgeBank:
		// address, err = registryInstance.BridgeBank(&auth)
		address = common.HexToAddress("0xF5c75E71240c7c346a299A2EE9532d1c5B384Fb4")
	default:
		panic("invalid target contract address")
	}

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return address, nil
}
