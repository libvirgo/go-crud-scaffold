package utils

import (
	"testing"
)

func TestVerifySig(t *testing.T) {
}

func TestSignHash(t *testing.T) {
	t.Log(SignHash([]byte("welcome login nonce:1698036540")))
}
