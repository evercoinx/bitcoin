package hash

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Hash160 hashes input data twice: first with SHA-256 and then with RIPEMD-160.
func Hash160(data []byte) []byte {
	d := sha256.Sum256(data)
	d2 := ripemd160.New()
	d2.Write(d[:])
	return d2.Sum(nil)
}

// Hash256 hashes input data two times with SHA-256.
func Hash256(data []byte) []byte {
	d := sha256.Sum256(data)
	d2 := sha256.Sum256(d[:])
	return d2[:]
}
