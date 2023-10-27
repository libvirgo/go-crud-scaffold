package utils

import "testing"

func TestString_IsWalletAddress(t *testing.T) {
	if String("0x0F3Bdd297496C14F2Be1085180151dad0f6e5B86").IsWalletAddress() == false {
		t.Fatal("0x0F3Bdd297496C14F2Be1085180151dad0f6e5B86 should be a wallet address")
	}
}
