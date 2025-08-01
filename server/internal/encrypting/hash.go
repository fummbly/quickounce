package encrypting

import (
	"crypto/sha1"
)

// Hashing function for creating picture urls
func Hash(b []byte) []byte {
	hasher := sha1.New()
	hasher.Write(b)
	return hasher.Sum(nil)
}
