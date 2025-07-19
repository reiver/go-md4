package md4_test

import (
	"testing"

	"github.com/reiver/go-md4"
)

func TestSum(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected [md4.Size]byte
	}{
		{
			Data: nil,
			Expected: [md4.Size]byte{0x31,0xd6,0xcf,0xe0,0xd1,0x6a,0xe9,0x31,0xb7,0x3c,0x59,0xd7,0xe0,0xc0,0x89,0xc0},
		},
		{
			Data: []byte(""),
			Expected: [md4.Size]byte{0x31,0xd6,0xcf,0xe0,0xd1,0x6a,0xe9,0x31,0xb7,0x3c,0x59,0xd7,0xe0,0xc0,0x89,0xc0},
		},



		{
			Data: []byte("Hello world!"),
			Expected: [md4.Size]byte{0x0d,0x7a,0x9d,0xb5,0xa3,0xbe,0xd4,0xae,0x57,0x38,0xee,0x6d,0x19,0x09,0x64,0x9c},
		},



		{
			Data: []byte("a"),
			Expected: [md4.Size]byte{0xbd,0xe5,0x2c,0xb3,0x1d,0xe3,0x3e,0x46,0x24,0x5e,0x05,0xfb,0xdb,0xd6,0xfb,0x24},
		},
		{
			Data: []byte("ab"),
			Expected: [md4.Size]byte{0xec,0x38,0x8d,0xd7,0x89,0x99,0xdf,0xc7,0xcf,0x46,0x32,0x46,0x56,0x93,0xb6,0xbf},
		},
		{
			Data: []byte("abc"),
			Expected: [md4.Size]byte{0xa4,0x48,0x01,0x7a,0xaf,0x21,0xd8,0x52,0x5f,0xc1,0x0a,0xe8,0x7a,0xa6,0x72,0x9d},
		},

		{
			Data: []byte("abcdefghijklmnopqrstuvwxyz"),
			Expected: [md4.Size]byte{0xd7,0x9e,0x1c,0x30,0x8a,0xa5,0xbb,0xcd,0xee,0xa8,0xed,0x63,0xdf,0x41,0x2d,0xa9},
		},



		{
			Data: []byte("message digest"),
			Expected: [md4.Size]byte{0xd9,0x13,0x0a,0x81,0x64,0x54,0x9f,0xe8,0x18,0x87,0x48,0x06,0xe1,0xc7,0x01,0x4b},
		},



		// This is longer than the MD4 block-size.
		{
			Data: []byte(
`B Boom Boom
B Boom Boom
Bees keep on flying 'round the room
M E Mer
M E Mer
Mother will help me if I ask her
R Ra Ran
R Ra Ran
I will run faster than they can
Ra Ran Rang
Ra Ran Rang
Put them together, it's boomerang!

Bing bong bang
Bing bong bang
Bing bong bang it's boomerang
Bing bong bang
Bing bong bang
Bing bong bang it's boomerang

A boomerang
A boomerang
What does it do?
What does it do?!
It comes back to you.
It comes back to you.
A boomerang
A boomerang
Bing bong bang it's boomerang

You throw a ball against the wall, and what does it do?
It comes back to you, like a boomerang.
Boomerang
Bing bong bang it's boomerang
We all have seen a trampoline, where you bounce so high,
And back down you go, like a boomerang
Boomerang
Bing bong bang it's boomerang

It's a funny looking thing
Something like an airplane wing
And when you throw it, ever so high
So high it seems to hide in the sky
It fools you with its special knack, of turning around and falling back

A boomerang
A boomerang
What does it do?
What does it do?!
It comes back to you.
It comes back to you.
A boomerang
A boomerang
Bing bong bang it's boomerang`),
			Expected: [md4.Size]byte{0x80,0x6e,0x5e,0x81,0x29,0x96,0x7d,0x95,0xce,0xcc,0xff,0x47,0x7d,0xf4,0xc0,0xe9},
		},
	}

	for testNumber, test := range tests {

		actual := md4.Sum(test.Data)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual md4-digest is not what was expected.", testNumber)
			t.Logf("EXPECTED: %X", expected)
			t.Logf("ACTUAL:   %X", actual)
			t.Logf("DATA: %X", test.Data)
			continue
		}
	}
}
