package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySig(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return false
	}
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(SignHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if fromAddr == recoveredAddr {
		return true
	}
	return false
}

func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
