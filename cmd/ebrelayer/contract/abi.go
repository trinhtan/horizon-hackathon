package contract

// -------------------------------------------------------
//    Contract Contains functionality for loading the
//				 smart contract
// -------------------------------------------------------

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/txs"
)

// File paths to Ethereum smart contract ABIs
const (
	BridgeBankABI     = "/generated/abi/BridgeBank/BridgeBank.abi"
	HarmonyBridgeABI  = "/generated/abi/HarmonyBridge/HarmonyBridge.abi"
	BridgeRegistryABI = "/generated/abi/BridgeRegistry/BridgeRegistry.abi"
)

// LoadABI loads a smart contract as an abi.ABI
func LoadABI(contractType txs.ContractRegistry) abi.ABI {
	var (
		_, b, _, _ = runtime.Caller(0)
		dir        = filepath.Dir(b)
	)

	var filePath string
	switch contractType {
	case txs.HarmonyBridge:
		filePath = HarmonyBridgeABI
	case txs.BridgeBank:
		filePath = BridgeBankABI
	case txs.BridgeRegistry:
		filePath = BridgeRegistryABI
	}

	// Read the file containing the contract's ABI
	contractRaw, err := ioutil.ReadFile(dir + filePath)
	if err != nil {
		panic(err)
	}

	// Convert the raw abi into a usable format
	contractABI, err := abi.JSON(strings.NewReader(string(contractRaw)))
	if err != nil {
		panic(err)
	}
	return contractABI
}
