package utils

import "regexp"

type String string

var walletRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func (s String) IsWalletAddress() bool {
	return walletRegex.MatchString(string(s))
}
