package txs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	oracle "github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/contract/generated/bindings/oracle"
	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/types"
)

const (
	// GasLimit the gas limit in Gwei used for transactions sent with TransactOpts
	GasLimit = uint64(3000000)
)

// // RelayProphecyClaimToEthereum relays the provided ProphecyClaim to HarmonyBridge contract on the Ethereum network
// func RelayProphecyClaimToEthereum(provider string, contractAddress common.Address, event types.Event,
// 	claim ProphecyClaim, key *ecdsa.PrivateKey) error {
// 	// Initialize client service, validator's tx auth, and target contract address
// 	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

// 	// Initialize HarmonyBridge instance
// 	fmt.Println("\nFetching HarmonyBridge contract...")
// 	harmonyBridgeInstance, err := harmonybridge.NewHarmonyBridge(target, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Send transaction
// 	fmt.Println("Sending new ProphecyClaim to HarmonyBridge...")
// 	// tx, err := harmonyBridgeInstance.NewUnlockClaim(auth,
// 	// 	claim.CosmosSender, claim.EthereumReceiver, claim.Symbol, claim.Amount)
// 	tx, err := harmonyBridgeInstance.NewUnlockClaim(auth,
// 		claim.EthereumReceiver, claim.EthereumReceiver, claim.EthereumReceiver, claim.Amount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("NewProphecyClaim tx hash:", tx.Hash().Hex())

// 	// Get the transaction receipt
// 	// receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// switch receipt.Status {
// 	// case 0:
// 	// 	fmt.Println("Tx Status: 0 - Failed")
// 	// case 1:
// 	// 	fmt.Println("Tx Status: 1 - Successful")
// 	// }
// 	return nil
// }

// RelayOracleClaimToEthereum relays the provided OracleClaim to Oracle contract on the Ethereum network
func RelayOracleClaimToEthereum(provider string, contractAddress common.Address, event types.Event,
	claim EthOracleClaim, key *ecdsa.PrivateKey) error {
	// Initialize client service, validator's tx auth, and target contract address
	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

	// Initialize Oracle instance
	fmt.Println("\nFetching Oracle contract...")
	oracleInstance, err := oracle.NewOracle(target, client)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	fmt.Println("Sending new OracleClaim to Oracle...")
	tx, err := oracleInstance.NewOracleClaim(auth, claim.UnlockID, claim.Message, claim.Signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NewOracleClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	// receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// switch receipt.Status {
	// case 0:
	// 	fmt.Println("Tx Status: 0 - Failed")
	// case 1:
	// 	fmt.Println("Tx Status: 1 - Successful")
	// }

	return nil
}

// initRelayConfig set up Ethereum client, validator's transaction auth, and the target contract's address
func initRelayConfig(provider string, registry common.Address, event types.Event, key *ecdsa.PrivateKey,
) (*ethclient.Client, *bind.TransactOpts, common.Address) {
	// Start Ethereum client
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	// Load the validator's address
	sender, err := LoadSender()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Set up TransactOpts auth's tx signature authorization
	transactOptsAuth := bind.NewKeyedTransactor(key)
	transactOptsAuth.Nonce = big.NewInt(int64(nonce))
	transactOptsAuth.Value = big.NewInt(0) // in wei
	transactOptsAuth.GasLimit = GasLimit
	transactOptsAuth.GasPrice = gasPrice

	var targetContract ContractRegistry
	switch event {
	// // ProphecyClaims are sent to the HarmonyBridge contract
	// case types.MsgBurn, types.MsgLock:
	// 	targetContract = HarmonyBridge
	// // OracleClaims are sent to the Oracle contract
	case types.EthLogNewUnlockClaim:
		targetContract = Oracle
	default:
		panic("invalid target contract address")
	}

	// Get the specific contract's address
	target, err := EthGetAddressFromBridgeRegistry(client, registry, targetContract)
	if err != nil {
		log.Fatal(err)
	}
	return client, transactOptsAuth, target
}
