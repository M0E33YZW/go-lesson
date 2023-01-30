package main

import "fmt"

func main() {
	n := 0b11100100
	// 8bitの塊の中から、下位1001のところだけ取り出す
	fmt.Printf("%b\n", n>>2&0b1111)

	/*
		n >> 2 -> 111001
				& 001111
				--------
				  001001

		0b1111	   : 15 | f
		0b11111111 : 255 | ff
	*/

	fmt.Printf("%b\n", 0b1_1111_1111)
	fmt.Printf("%b\n", 0x1ff)

	fmt.Printf("%b\n", 0b11_1111_1111)
	fmt.Printf("%b\n", 0x3ff)

	fmt.Printf("%b\n", 0b111_1111_1111)
	fmt.Printf("%b\n", 0x7ff)

	fmt.Printf("%b\n", 0b1111_1111_1111)
	fmt.Printf("%b\n", 0xfff)
}
