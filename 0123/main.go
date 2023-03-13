package main

import (
	"fmt"
	"unsafe"
	// "strings"
)

type width8 interface {
	int8 | uint8
}

func BinStr[T width8](n T) string {
	/* up := unsafe.Pointer(&n) // Pointer型
	up8 := (*uint8)(up)         // Pointer型
	u8 := *up8                  // uint8型 */

	u8 := *(*uint8)(unsafe.Pointer(&n))
	return fmt.Sprintf("%08b", u8)
}

func main() {
	binStr := BinStr(int8(-50))
	fmt.Println(binStr)
}

/* func ToString(a []int) string {
	sb := strings.Builder{}

	sb.WriteRune('[')
	// sb.WriteString("[") でも同じことが可能だが、コスト高
	if len(a) > 0 {
		sb.WriteString(fmt.Sprintf("%d", a[0]))
		for i := 1; i < len(a); i++ {
			sb.WriteString(fmt.Sprintf(", %d", a[i])) // ,で繋げて表示
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func main() {
	// a := 50
	// fmt.Printf("%08b, %04x, %d\n", a, a, a)

	// b := -50
	// fmt.Printf("%08b, %04x, %d\n", b, b, b)
	s := ToString([]int{1, 2})
	fmt.Println(s)
} */
