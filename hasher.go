package md4

import (
	"encoding/binary"
	"hash"

	"github.com/reiver/go-erorr"
)

// hasher is used to calculate the MD4 hash-function.
type hasher struct {
	state  [4]uint32 // this is defined in IETF RFC-1320
	buffer [BlockSize]byte
	index  int // this is the index into buffer
	length uint64 // this is the total amount written into the hasher
}

var _ hash.Hash = &hasher{}

// Reset resets hasher to the initial state,
// as per how it is defined in IETF RFC-1320.
//
// https://datatracker.ietf.org/doc/html/rfc1320
func (receiver *hasher) Reset() {
	if nil == receiver {
		return
	}

	// Load magic initialization constants.
	//
	// These hexadecimal constants are taken from IETF RFC-1320.
	// And, are hard-coded there like this, too.
	//
	// https://datatracker.ietf.org/doc/html/rfc1320
	receiver.state[0] = 0x67452301
	receiver.state[1] = 0xEFCDAB89
	receiver.state[2] = 0x98BADCFE
	receiver.state[3] = 0x10325476

	receiver.index = 0
	receiver.length = 0
}

// Size returns the size of the MD4 digest measured in byte.
//
// Size will always return 16.
//
// So, the size of the MD4 digest is 16 bytes (i.e., 128 bits).
func (hasher) Size() int {
	return Size
}

// BlockSize returns the size of the MD4 block measured in byte.
//
// Size will always return 64.
//
// So, the size of the MD4 block is 64 bytes (i.e., 512 bits).
func (hasher) BlockSize() int {
	return BlockSize
}

// Write puts in more data into the ongoing MD4 hasher.
//
// Write makes hasher match the io.Writer interface.
func (receiver *hasher) Write(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	n = len(p)
	receiver.length += uint64(n)

	if receiver.index > 0 {
		numCopied := copy(receiver.buffer[receiver.index:], p)
		receiver.index += numCopied
		if receiver.index == BlockSize {
			receiver.processBlock(receiver.buffer[:])
			receiver.index = 0
		}
		p = p[numCopied:]
	}

	for len(p) >= BlockSize {

		// .processBlock() will only process the first 64 bytes of the slise,
		// so do NOT need to make the slice passed to .processBlock() be
		// exactly 64 bytes.
		receiver.processBlock(p)
		p = p[BlockSize:]
	}

	if len(p) > 0 {
		receiver.index = copy(receiver.buffer[:], p)
	}

	return
}

// Sum appends the MD4 digest of the data that had been written to the hasher so far.
//
// Sum appends the the digest to `p` (rather than just returning the digest).
//
// To have Sum return the digest, pass nil to Sun.
// For example:
//
//	digest := hasher.Sum(nil)
func (receiver *hasher) Sum(p []byte) []byte {

	// Make a copy of the hasher, so that when sum() writes to the receiver
	// as part of calculating the (current) digest, it doesn't affect the
	// original receiver. And, the original receiver can have more data written to it.
	clone := *receiver

	digest := clone.sum()
	return append(p, digest[:]...)
}

// sum returns the MD4 digest (checksum) of the data already written to the hasher.
//
// Note that sum mutates the receiver.
func (receiver *hasher) sum() [Size]byte {
	if nil == receiver {
		return nothing
	}

	length := receiver.length

	// Note that we take advantage of the fact that `padding` is
	// initially initialized to all zeros.
	var padding [BlockSize]byte
	padding[0] = 0x80
	{
		var padLength int
		if (length % BlockSize) < 56 {
			padLength = 56 - int(length % BlockSize)

		} else {
			padLength = (BlockSize+56) - int(length % BlockSize)
		}

		receiver.Write(padding[0:padLength])
	}
	{
		length <<= 3
		for i := uint(0); i < 8; i++ {
			padding[i] = byte(length >> (8 * i))
		}

		receiver.Write(padding[0:8])
	}

	// This should never happen.
	if 0 != receiver.index {
		panic(erorr.Errorf("md4: recevier.index != 0 in .sum() — recevier.index == %d", receiver.index))
	}

	var digest [Size]byte
	{
		binary.LittleEndian.PutUint32(digest[0:], receiver.state[0])
		binary.LittleEndian.PutUint32(digest[4:], receiver.state[1])
		binary.LittleEndian.PutUint32(digest[8:], receiver.state[2])
		binary.LittleEndian.PutUint32(digest[12:], receiver.state[3])
	}

	return digest
}

// processBlock processes a single 64-byte (i.e., 512-bit) block by using the MD4 "MD4Transform" function.
func (receiver *hasher) processBlock(data []byte) {
	if nil == receiver {
		return
	}

	data = data[:BlockSize]

	a, b, c, d := md4transform(receiver.state, data)

	receiver.state[0] += a
	receiver.state[1] += b
	receiver.state[2] += c
	receiver.state[3] += d
}
