package md4

// Size is the size of an MD4 digest (checksum) measured in bytes.
//
// The size of an MD4 digest 16 bytes (i.e., 128 bits).
// For example:
//
//	//          1    2    3    4    5    6    7    8    9   10   11   12   13   14   15   16
//	[16]byte{0x31,0xd6,0xcf,0xe0,0xd1,0x6a,0xe9,0x31,0xb7,0x3c,0x59,0xd7,0xe0,0xc0,0x89,0xc0}
//
// Note that if you are representing an MD4 digest in hexadecimal, then it will double the number of characters to 32 bytes.
// For example:
//
//	"31d6cfe0d16ae931b73c59d7e0c089c0"
//
// Or, if you include the "0x" prefix in front of the hexadecimal representation, then it will be 32+2=34 bytes long:
//
//	0x31d6cfe0d16ae931b73c59d7e0c089c0
//
// Size represents the length of the MD4 digest in a binary representation.
const Size = 16
