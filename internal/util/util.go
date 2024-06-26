package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strconv"
	"strings"
)

func Sha3256Hash(bytes [][]byte) (output []byte) {
	hasher := sha3.New256()
	for _, b := range bytes {
		hasher.Write(b)
	}
	return hasher.Sum([]byte{})
}

// ParseHex Convenience function to deal with 0x at the beginning of hex strings
func ParseHex(hexStr string) ([]byte, error) {
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = hexStr[2:]
	}
	return hex.DecodeString(hexStr)
}

func BytesToHex(bytes []byte) string {
	return "0x" + hex.EncodeToString(bytes)
}

func StrToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StrToBigInt(val string) (num *big.Int, err error) {
	num = &big.Int{}
	_, ok := num.SetString(val, 10)
	if !ok {
		return nil, fmt.Errorf("num %s is not an integer", val)
	}
	return num, nil
}

// PrettyJson a simple pretty print for JSON examples
func PrettyJson(x any) string {
	out := strings.Builder{}
	enc := json.NewEncoder(&out)
	enc.SetIndent("", "  ")
	err := enc.Encode(x)
	if err != nil {
		return ""
	}
	return out.String()
}
