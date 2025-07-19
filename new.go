package md4

import (
	"hash"
)

// New returns a hash.Hash that calculates an MD4 digest (checksum) using the MD4 cryptographic hash-function.
func New() hash.Hash {
	var h hash.Hash = newHasher()
	return h
}

func newHasher() *hasher {
	var h hasher
	h.Reset()
	return &h
}
