package main

import (
	"fmt"
	"unsafe"
)

type width8 interface {
	~uint8 | ~int8 // ~を付けると、int8又はint8を継承した値を扱える（汎用的）
}

func BinStr[T width8](n T) string {
	/* up := unsafe.Pointer(&n) //Pointer 型
	up8 := (*uint8)(up) //Pointer 型
	u8 := *up8 // uint8型
	以上3行は、以下1行で書ける */
	u8 := *(*uint8)(unsafe.Pointer(&n))
	return fmt.Sprintf("%08b", u8)
}

// 符号付
func NewInt8(n int) int8 {
	return int8(n)
}

// 符号なし（unsigned）
func NewUint8(n int) uint8 {
	return uint8(n)
}

func main() {
	/*
		fmt.Println("Hello world")
		fmt.Printf("私は%d歳です。名前は%sです。お兄さんは%d歳です。\n", 3, "山田", 10)
		fmt.Printf("私は%v歳です。名前は%vです。お兄さんは%v歳です。\n", 3, "山田", 10)
		fmt.Printf("私は%x歳です。名前は%sです。お兄さんは%x歳です。\n", 3, "山田", 10)
		fmt.Printf("私は%b歳です。名前は%sです。お兄さんは%b歳です。\n", 3, "山田", 10)

		// n := 11
		var n int
		n = 11
		fmt.Printf("10進: %d, 16進: %x, 2進: %b \n", n, n, n)
	*/

	/*
		// n := 0xa
		n := 0b1010
		fmt.Printf("%04b | 0001 => %04b\n", n, n|0b0001)
		fmt.Printf("%04b & 0001 => %04b\n", n, n&0b0001)
		fmt.Println()
	*/

	/* or i := 0; i < 16; i++ {
		fmt.Printf("%04b & 1001 => %04b\n", i, i&0b1010) // AND 特定ビット(0のビット)をオフにする
		// fmt.Printf("%04b | 1110 => %04b\n", i, i|0b1010) // OR 特定ビット(1のビット)をオンにする
		// fmt.Printf("%04b ^ 1111 => %04b\n", i, i^0b1010) // XOR, EOR 特定ビット(1のビット)を反転する
	}

	for i := 0; i < 16; i++ {
		fmt.Printf("%04b | 1110 => %04b\n", i, i|0b1010) // OR 特定ビット(1のビット)をオンにする
	}

	for i := 0; i < 16; i++ {
		fmt.Printf("%04b ^ 1111 => %04b\n", i, i^0b1010) // XOR, EOR 特定ビット(1のビット)を反転する
	} */

	// n := 0b0111
	// fmt.Printf("%04b (%d) << 1 = %04b (%d) \n", n, n, n<<1, n<<1)
	// fmt.Printf("%04b (%d) >> 1 = %04b (%d) \n", n, n, n>>1, n>>1)

	n := NewUint8(0b11111010)
	fmt.Printf("%s (%d) %s(%d) \n", BinStr(n), n, BinStr(n>>1), n>>1)
	
	n2 := NewInt8(0b11111010) // 算術シフト
	fmt.Printf("%s (%d) %s(%d) \n", BinStr(n2), n2, BinStr(n2>>1), n2>>1)
}

/*
	論理積（AND）
	0 & 0 = 0	F & F = F
	0 & 1 = 0	F & F = F
	1 & 0 = 0	T & F = F
	1 & 1 = 1	T & T = T

	論理和（OR）
	0 | 0 = 0	F | F = F
	0 | 1 = 1	F | T = T
	1 | 0 = 1	T | F = T
	1 | 1 = 1	T | T = T

	排他的論理和（XOR）
	0 ^ 0 = 0	F ^ F = F
	0 ^ 1 = 1	F ^ T = T
	1 ^ 0 = 1	T ^ F = T
	1 ^ 1 = 0	T ^ T = F
*/
