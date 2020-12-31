package txs

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/trinhtan/horizon-hackathon/cmd/ebrelayer/types"
)

const (
	nullAddress = "0x0000000000000000000000000000000000000000"
)

// UnlockClaimToSignedOracleClaim packages and signs a unlock claim's data, returning a new oracle claim
func UnlockClaimToSignedOracleClaim(event types.EthLogNewUnlockClaimEvent, key *ecdsa.PrivateKey) (EthOracleClaim, error) {
	oracleClaim := EthOracleClaim{}

	// Generate a hashed claim message which contains ProphecyClaim's data
	fmt.Println("Generating unique message for UnlockClaim", event.UnlockID)
	message := GenerateClaimMessage(event)

	// Sign the message using the validator's private key
	fmt.Println("Signing message...")
	signature, err := SignClaim(PrefixMsg(message), key)
	if err != nil {
		return oracleClaim, err
	}
	fmt.Println("Signature generated:", hexutil.Encode(signature))

	oracleClaim.UnlockID = event.UnlockID
	var message32 [32]byte
	copy(message32[:], message)
	oracleClaim.Message = message32
	oracleClaim.Signature = signature
	return oracleClaim, nil
}

// isZeroAddress checks an Ethereum address and returns a bool which indicates if it is the null address
func isZeroAddress(address common.Address) bool {
	return address == common.HexToAddress(nullAddress)
}
