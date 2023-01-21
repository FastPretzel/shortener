package usecase

import (
	"hash/crc64"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const alphLen = 63

func hashing(s string) string {
	// Hashing given string with CRC-64
	table := crc64.MakeTable(crc64.ISO)
	hash := crc64.Checksum([]byte(s), table)

	charArray := make([]uint8, 0, 10)

	// Converting uint64 to {alphabetLen} numeral system
	// with dropping last symbols
	for mod := uint64(0); hash != 0 && len(charArray) < 10; {
		mod = hash % alphLen
		hash /= alphLen
		charArray = append(charArray, alphabet[mod])
	}

	// prepend zero symbols if result array len < 10
	// and convert to string
	return strings.Repeat(string(alphabet[0]), 10-len(charArray)) + string(charArray)
}
