package md4

import (
	"encoding/binary"
)

// md4transform is an implementation of "MD4Transform" from IETF RFC-1320.
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
// md4transform implements the "MD4 basic transformation".
//
// Although NOTE that this returns `a`, `b`, `c`, and `d` rather than
// adding them to the state.
//
// (That is where this implementation of md4transform differs from IETF RFC-1320.)
//
// That adding is done outside of this function.
func md4transform(state [4]uint32, data []byte) (a uint32, b uint32, c uint32, d uint32) {

	var x [16]uint32

	for i := 0; i < 16; i++ {
		x[i] = binary.LittleEndian.Uint32(data[4*i:])
	}

	a, b, c, d = state[0], state[1], state[2], state[3]

	// round 1
	a = ff(a, b, c, d, x[ 0], s11)
	d = ff(d, a, b, c, x[ 1], s12)
	c = ff(c, d, a, b, x[ 2], s13)
	b = ff(b, c, d, a, x[ 3], s14)
	a = ff(a, b, c, d, x[ 4], s11)
	d = ff(d, a, b, c, x[ 5], s12)
	c = ff(c, d, a, b, x[ 6], s13)
	b = ff(b, c, d, a, x[ 7], s14)
	a = ff(a, b, c, d, x[ 8], s11)
	d = ff(d, a, b, c, x[ 9], s12)
	c = ff(c, d, a, b, x[10], s13)
	b = ff(b, c, d, a, x[11], s14)
	a = ff(a, b, c, d, x[12], s11)
	d = ff(d, a, b, c, x[13], s12)
	c = ff(c, d, a, b, x[14], s13)
	b = ff(b, c, d, a, x[15], s14)

	// round 2
	a = gg(a, b, c, d, x[ 0], s21)
	d = gg(d, a, b, c, x[ 4], s22)
	c = gg(c, d, a, b, x[ 8], s23)
	b = gg(b, c, d, a, x[12], s24)
	a = gg(a, b, c, d, x[ 1], s21)
	d = gg(d, a, b, c, x[ 5], s22)
	c = gg(c, d, a, b, x[ 9], s23)
	b = gg(b, c, d, a, x[13], s24)
	a = gg(a, b, c, d, x[ 2], s21)
	d = gg(d, a, b, c, x[ 6], s22)
	c = gg(c, d, a, b, x[10], s23)
	b = gg(b, c, d, a, x[14], s24)
	a = gg(a, b, c, d, x[ 3], s21)
	d = gg(d, a, b, c, x[ 7], s22)
	c = gg(c, d, a, b, x[11], s23)
	b = gg(b, c, d, a, x[15], s24)

	// round 3
	a = hh(a, b, c, d, x[ 0], s31)
	d = hh(d, a, b, c, x[ 8], s32)
	c = hh(c, d, a, b, x[ 4], s33)
	b = hh(b, c, d, a, x[12], s34)
	a = hh(a, b, c, d, x[ 2], s31)
	d = hh(d, a, b, c, x[10], s32)
	c = hh(c, d, a, b, x[ 6], s33)
	b = hh(b, c, d, a, x[14], s34)
	a = hh(a, b, c, d, x[ 1], s31)
	d = hh(d, a, b, c, x[ 9], s32)
	c = hh(c, d, a, b, x[ 5], s33)
	b = hh(b, c, d, a, x[13], s34)
	a = hh(a, b, c, d, x[ 3], s31)
	d = hh(d, a, b, c, x[11], s32)
	c = hh(c, d, a, b, x[ 7], s33)
	b = hh(b, c, d, a, x[15], s34)

	return a, b, c, d
}
