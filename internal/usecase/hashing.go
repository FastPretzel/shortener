package usecase

import (
	"hash/crc64"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const alphLen = 63

func hashing(s string) string {
	table := crc64.MakeTable(crc64.ISO)
	hash := crc64.Checksum([]byte(s), table)

	charArray := make([]uint8, 0, 10)

	for mod := uint64(0); hash != 0 && len(charArray) < 10; {
		mod = hash % alphLen
		hash /= alphLen
		charArray = append(charArray, alphabet[mod])
	}

	return strings.Repeat(string(alphabet[0]), 10-len(charArray)) + string(charArray)
}
