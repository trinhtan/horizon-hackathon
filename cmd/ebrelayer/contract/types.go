package contract

// BridgeContract is an enum containing supported contract names
type BridgeContract int

const (
	// BridgeRegistry registers deployed addresses of the other contracts
	BridgeRegistry BridgeContract = iota + 1
	// Valset manages the validator set and signature verification
	Valset
	// Oracle enables validators to make OracleClaims and processes UnlockClaims
	Oracle
	// HarmonyBridge enables validators to make UnlockClaims
	HarmonyBridge
	// EthereumBridge enables validators to make UnlockClaims
	EthereumBridge
	// BridgeBank manages protocol assets on both Ethereum and Harmony
	BridgeBank
)

// BridgeContractToString returns the string associated with a BridgeContract
var BridgeContractToString = [...]string{"BridgeRegistry", "Valset", "Oracle", "HarmonyBridge", "EthereumBridge", "BridgeBank"}

// String returns the BridgeContract as a string
func (d BridgeContract) String() string {
	return BridgeContractToString[d-1]
}

// BridgeContracts is a slice of BridgeContract
type BridgeContracts []BridgeContract

// LoadEthereumBridgeContracts loads populated slice of BridgeContract
func LoadEthereumBridgeContracts() BridgeContracts {
	return BridgeContracts{
		BridgeRegistry,
		Valset,
		Oracle,
		HarmonyBridge,
		BridgeBank,
	}
}

// LoadHarmonyBridgeContracts loads populated slice of BridgeContract
func LoadHarmonyBridgeContracts() BridgeContracts {
	return BridgeContracts{
		BridgeRegistry,
		Valset,
		Oracle,
		EthereumBridge,
		BridgeBank,
	}
}
