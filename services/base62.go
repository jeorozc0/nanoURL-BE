package services

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/google/uuid"
)


const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// UUIDToShortID converts a UUID to a 4-character base62 ID
func UUIDToShortID(u uuid.UUID) string {
	// Hash the UUID to get a more uniform distribution
	hash := sha256.Sum256(u[:])

	// Take the first 4 bytes of the hash and convert to uint32
	num := binary.BigEndian.Uint32(hash[:4])

	// Convert to base62
	return encodeBase62(big.NewInt(int64(num)))[:4]
}

// encodeBase62 encodes a big.Int to a base62 string
func encodeBase62(n *big.Int) string {
	if n.Sign() == 0 {
		return string(base62Chars[0])
	}

	chars := []byte{}
	base := big.NewInt(62)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for n.Cmp(zero) > 0 {
		n.DivMod(n, base, mod)
		chars = append(chars, base62Chars[mod.Int64()])
	}

	// Reverse the characters
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
