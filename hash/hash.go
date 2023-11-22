package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hash converts a string to a base64 encoded hash.
func Hashe(input string) string {
	// Implement your hashing logic here using SHA-256.
	hashFunction := sha256.New()
	hashFunction.Write([]byte(input))
	hashValue := hashFunction.Sum(nil)

	// Truncate the hash value and convert to base64 encoding
	truncatedHash := hashValue[:8] // Adjust the length as needed
	encodedHash := base64.URLEncoding.EncodeToString(truncatedHash)
	return encodedHash
}
