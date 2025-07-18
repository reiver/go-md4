package md4

// f is an implementation of "F" that is defined in section 3.4 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
//	F(X,Y,Z) = XY v not(X) Z
//
// Drawn from:
//
//	We first define three auxiliary functions that each take as input
//	three 32-bit words and produce as output one 32-bit word.
//
//	F(X,Y,Z) = XY v not(X) Z
//	G(X,Y,Z) = XY v XZ v YZ
//	H(X,Y,Z) = X xor Y xor Z
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func f(x, y, z uint32) uint32 {
        return (x & y) | (^x & z)
}

// g is an implementation of "G" that is defined in section 3.4 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
//	G(X,Y,Z) = XY v XZ v YZ
//
// Drawn from:
//
//	We first define three auxiliary functions that each take as input
//	three 32-bit words and produce as output one 32-bit word.
//
//	F(X,Y,Z) = XY v not(X) Z
//	G(X,Y,Z) = XY v XZ v YZ
//	H(X,Y,Z) = X xor Y xor Z
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func g(x, y, z uint32) uint32 {
        return (x & y) | (x & z) | (y & z)
}

// h is an implementation of "H" that is defined in section 3.4 of IETF RFC-1320
//
// https://datatracker.ietf.org/doc/html/rfc1320
//
//	H(X,Y,Z) = X xor Y xor Z
//
// Drawn from:
//
//	We first define three auxiliary functions that each take as input
//	three 32-bit words and produce as output one 32-bit word.
//
//	F(X,Y,Z) = XY v not(X) Z
//	G(X,Y,Z) = XY v XZ v YZ
//	H(X,Y,Z) = X xor Y xor Z
//
// IETF RFC-1320 is what gives this function an undescriptive name.
// I have kept it the same here to try to make it easier to match
// this source-code with what is in IETF RFC-1320.
func h(x, y, z uint32) uint32 {
        return x ^ y ^ z
}
