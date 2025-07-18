package md4

// rotateLeft is an implementation of "ROTATE_LEFT" that is defined in section A.3 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
func rotateLeft(value uint32, shift uint) uint32 {
	return (value << shift) | (value >> (32 - shift))
}

// ff is an implementation of "FF" that is defined in section A.3 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func ff(a, b, c, d, x uint32, s uint) uint32 {
	return rotateLeft(a+f(b,c,d)+x, s)
}

// gg is an implementation of "GG" that is defined in section A.3 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func gg(a, b, c, d, x uint32, s uint) uint32 {
	return rotateLeft(a+g(b,c,d)+x+0x5A827999, s)
}

// hh is an implementation of "HH" that is defined in section A.3 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func hh(a, b, c, d, x uint32, s uint) uint32 {
	return rotateLeft(a+h(b,c,d)+x+0x6ED9EBA1, s)
}
