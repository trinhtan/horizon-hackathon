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
	// LogLock is for Ethereum event LogLock
	LogLock
	// LogNewUnlockClaim is an Ethereum event named 'LogNewUnlockClaim'
	LogNewUnlockClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "LogLock", "LogNewUnlockClaim"}[d]
}

// EthereumLogLockEvent struct is used by LogLock
type EthereumLogLockEvent struct {
	EthereumChainID     *big.Int
	BridgeBankAddress   common.Address
	ID                  [32]byte
	EthereumSender      common.Address
	HarmonyReceiver     common.Address
	EthereumToken       common.Address
	HarmonyToken        common.Address
	EthereumTokenAmount *big.Int
	HarmonyTokenAmount  *big.Int
	Nonce               *big.Int
}

// String implements fmt.Stringer
func (e EthereumLogLockEvent) String() string {
	return fmt.Sprintf("\nChain ID: %v\nBridge contract address: %v\nEthereum Token: %v\nHarmony Token: %v\nEthereum Sender: %v\nHarmony Recipient: %v\nEthereum Token Amount: %v\nHarmony Token Amount: %v\nNonce: %v\n",
		e.EthereumChainID, e.BridgeBankAddress.Hex(), e.EthereumToken.Hex(), e.HarmonyToken.Hex(), e.EthereumSender.Hex(),
		e.HarmonyReceiver.Hex(), e.EthereumTokenAmount, e.HarmonyTokenAmount, e.Nonce)
}

// EthereumLogNewUnlockClaimEvent struct which represents a LogNewUnlockClaim event
type EthereumLogNewUnlockClaimEvent struct {
	UnlockID         *big.Int
	HarmonySender    common.Address
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Amount           *big.Int
}

// NewEthereumLogNewUnlockClaimEvent creates a new EthereumLogNewUnlockClaimEvent
func NewEthereumLogNewUnlockClaimEvent(harmonySender common.Address, unlockID, amount *big.Int, ethereumReceiver,
	validatorAddress, tokenAddress common.Address, claimType uint8) EthereumLogNewUnlockClaimEvent {
	return EthereumLogNewUnlockClaimEvent{
		HarmonySender:    harmonySender,
		UnlockID:         unlockID,
		Amount:           amount,
		EthereumReceiver: ethereumReceiver,
		ValidatorAddress: validatorAddress,
		TokenAddress:     tokenAddress,
	}
}

// String implements fmt.Stringer
func (p EthereumLogNewUnlockClaimEvent) String() string {
	return fmt.Sprintf("\nUnlocl ID: %v\nHarmony Sender: %v\n"+
		"Ethereum Receiver: %v\nEthereum Token: %v\nAmount: %v\nValidator Address: %v\n\n",
		p.UnlockID, p.HarmonySender.Hex(), p.EthereumReceiver.Hex(),
		p.TokenAddress.Hex(), p.Amount, p.ValidatorAddress.Hex())
}
