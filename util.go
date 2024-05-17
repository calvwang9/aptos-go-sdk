package aptos

import (
	"github.com/aptos-labs/aptos-go-sdk/internal/util"
)

// -- Note these are copied from internal/util/util.go to prevent package loops, but still allow devs to use it

// ParseHex Convenience function to deal with 0x at the beginning of hex strings
func ParseHex(hexStr string) ([]byte, error) {
	// This had to be redefined separately to get around a package loop
	return util.ParseHex(hexStr)
}

// SHA3_256Hash takes a hash of the given sets of bytes
func SHA3_256Hash(bytes [][]byte) (output []byte) {
	return util.SHA3_256Hash(bytes)
}