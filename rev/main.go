// FE PMB No.6
package main

import (
	"fmt"
)

func rev(bin byte) byte {
	var rbyte byte = bin
	var r byte = 0

	for i := 0; i < 8; i++ {
		r = (r << 1) | (rbyte & 1)
		rbyte >>= 1

	}
	return r
}

func main() {
	fmt.Printf("%08b\n", rev(byte(0b01001011)))
}
