package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Event enum containing supported chain events
type Event byte

const (
	// Unsupported is an invalid Cosmos or Ethereum event
	Unsupported Event = iota
	// MsgBurn is a Cosmos msg of type MsgBurn
	MsgBurn
	// MsgLock is a Cosmos msg of type MsgLock
	MsgLock
	// LogLock is for Ethereum event LogLock
	LogLock
	// LogNewUnlockClaim is an Ethereum event named 'LogNewProphecyClaim'
	LogNewUnlockClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "burn", "lock", "LogLock", "LogNewProphecyClaim", "LogNewUnlockClaim"}[d]
}

// EthereumEvent struct is used by LogLock
type EthereumEvent struct {
	EthereumChainID       *big.Int
	BridgeContractAddress common.Address
	ID                    [32]byte
	EthereumSender        common.Address
	HarmonyReceiver       common.Address
	EthereumToken         common.Address
	HarmonyToken          common.Address
	Value                 *big.Int
	Nonce                 *big.Int
}

// String implements fmt.Stringer
func (e EthereumEvent) String() string {
	return fmt.Sprintf("\nChain ID: %v\nBridge contract address: %v\nEthereum Token: %v\nHarmony Token "+
		"contract address: %v\nSender: %v\nRecipient: %v\nValue: %v\nNonce: %v",
		e.EthereumChainID, e.BridgeContractAddress.Hex(), e.EthereumToken.Hex(), e.HarmonyToken.Hex(), e.EthereumSender.Hex(),
		e.HarmonyReceiver.Hex(), e.Value, e.Nonce)
}

// UnlockClaimEvent struct which represents a LogNewUnlockClaim event
type UnlockClaimEvent struct {
	HarmonySender    common.Address
	ProphecyID       *big.Int
	Amount           *big.Int
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
}

// NewUnlockClaimEvent creates a new UnlockClaimEvent
func NewUnlockClaimEvent(harmonySender common.Address, prophecyID, amount *big.Int, ethereumReceiver,
	validatorAddress, tokenAddress common.Address, claimType uint8) UnlockClaimEvent {
	return UnlockClaimEvent{
		HarmonySender:    harmonySender,
		ProphecyID:       prophecyID,
		Amount:           amount,
		EthereumReceiver: ethereumReceiver,
		ValidatorAddress: validatorAddress,
		TokenAddress:     tokenAddress,
	}
}

// String implements fmt.Stringer
func (p UnlockClaimEvent) String() string {
	return fmt.Sprintf("\nProphecy ID: %v\nHarmony Sender: %v\n"+
		"Recipient: %v\nEthereum Token: %v\nAmount: %v\nValidator: %v\n\n",
		p.ProphecyID, p.HarmonySender.Hex(), p.EthereumReceiver.Hex(),
		p.TokenAddress.Hex(), p.Amount, p.ValidatorAddress.Hex())
}

// CosmosMsg contains data from MsgBurn and MsgLock events
type CosmosMsg struct {
	ClaimType        Event
	CosmosSender     []byte
	EthereumReceiver common.Address
	Symbol           string
	Amount           *big.Int
}

// NewCosmosMsg creates a new CosmosMsg
func NewCosmosMsg(claimType Event, cosmosSender []byte, ethereumReceiver common.Address, symbol string,
	amount *big.Int) CosmosMsg {
	return CosmosMsg{
		ClaimType:        claimType,
		CosmosSender:     cosmosSender,
		EthereumReceiver: ethereumReceiver,
		Symbol:           symbol,
		Amount:           amount,
	}
}

// String implements fmt.Stringer
func (c CosmosMsg) String() string {
	if c.ClaimType == MsgLock {
		return fmt.Sprintf("\nClaim Type: %v\nCosmos Sender: %v\nEthereum Recipient: %v"+
			"\nSymbol: %v\nAmount: %v\n",
			c.ClaimType.String(), string(c.CosmosSender), c.EthereumReceiver.Hex(), c.Symbol, c.Amount)
	}
	return fmt.Sprintf("\nClaim Type: %v\nCosmos Sender: %v\nEthereum Recipient: %v"+
		"\nSymbol: %v\nAmount: %v\n",
		c.ClaimType.String(), string(c.CosmosSender), c.EthereumReceiver.Hex(), c.Symbol, c.Amount)
}

// CosmosMsgAttributeKey enum containing supported attribute keys
type CosmosMsgAttributeKey int

const (
	// UnsupportedAttributeKey unsupported attribute key
	UnsupportedAttributeKey CosmosMsgAttributeKey = iota
	// CosmosSender sender's address on Cosmos network
	CosmosSender
	// EthereumReceiver receiver's address on Ethereum network
	EthereumReceiver
	// Amount is coin's value
	Amount
	// Symbol is the coin type
	Symbol
)

// String returns the event type as a string
func (d CosmosMsgAttributeKey) String() string {
	return [...]string{"unsupported", "cosmos_sender", "ethereum_receiver", "amount", "symbol"}[d]
}
