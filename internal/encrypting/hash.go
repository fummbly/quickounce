package encrypting

import (
	"crypto/sha1"
)

func Hash(b []byte) []byte {
	hasher := sha1.New()
	hasher.Write(b)
	return hasher.Sum(nil)
}
