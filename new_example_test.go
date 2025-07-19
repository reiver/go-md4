package md4_test

import (
	"fmt"
	"io"

	"github.com/reiver/go-md4"
)

func ExampleNew() {

	hasher := md4.New()

	// “Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.”
	// ⸺ Rumi

	io.WriteString(hasher, "“")
	io.WriteString(hasher, "Yesterday I was clever, so I wanted to change the world.")
	io.WriteString(hasher, " ")
	io.WriteString(hasher, "Today I am wise, so I am changing myself.")
	io.WriteString(hasher, "”")
	io.WriteString(hasher, "\n")
	io.WriteString(hasher, "”")
	io.WriteString(hasher, "⸺")
	io.WriteString(hasher, " ")
	io.WriteString(hasher, "Rumi")

	digest := hasher.Sum(nil)

	fmt.Printf("MD4-DIGEST: 0x%032X\n", digest)

	// Output:
	// MD4-DIGEST: 0xA2D6BDF512B93A4DAD4556D8F9495E5B
}
