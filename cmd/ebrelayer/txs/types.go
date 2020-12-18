package txs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// OracleClaim contains data required to make an OracleClaim
type OracleClaim struct {
	UnlockID  *big.Int
	Message   [32]byte
	Signature []byte
}

// UnlockClaim contains data required to make an ProphecyClaim
type UnlockClaim struct {
	HarmonySender    common.Address
	EthereumReceiver common.Address
	Token            common.Address
	Amount           *big.Int
}
