package md4_test

import (
	"fmt"

	"github.com/reiver/go-md4"
)

func ExampleSum() {

	var data []byte = []byte("Hello world!")

	digest := md4.Sum(data)

	fmt.Printf("MD4-DIGEST: 0x%032X\n", digest)

	// Output:
	// MD4-DIGEST: 0x0D7A9DB5A3BED4AE5738EE6D1909649C
}
