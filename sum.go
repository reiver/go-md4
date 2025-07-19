package md4

// Sum returns the MD4 digest (checksum) of `data` calculated using the MD4 hash-function.
func Sum(data []byte) [Size]byte {
	var h hasher
        h.Reset()
	h.Write(data)
	return h.sum()
}
