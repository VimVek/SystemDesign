package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hashe converts a string to a base64 encoded hash.
func Hashe(input string) string {

	hashFunction := sha256.New()
	hashFunction.Write([]byte(input))
	hashValue := hashFunction.Sum(nil)

	// Truncate the hash value and convert to base64 encoding
	truncatedHash := hashValue[:8]
	encodedHash := base64.URLEncoding.EncodeToString(truncatedHash)
	return encodedHash
}
