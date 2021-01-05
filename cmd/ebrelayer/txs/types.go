package txs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// EthOracleClaim contains data required to make an EthOracleClaim
type EthOracleClaim struct {
	UnlockID  *big.Int
	Message   [32]byte
	Signature []byte
}

// EthUnlockClaim contains data required to make an ProphecyClaim
type EthUnlockClaim struct {
	HarmonySender    common.Address
	EthereumReceiver common.Address
	Token            common.Address
	Amount           *big.Int
}
