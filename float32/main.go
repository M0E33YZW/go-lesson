package main

import (
	"fmt"
	"math"
	"os"
	// "reflect"
	"strconv"
	"unsafe"
)

func Float32Internal(f float32) (sign, exponent, fraction uint32) {
	decimal := f
	fmt.Printf("decimal  : %f\n", decimal)
	binary := *(*uint32)(unsafe.Pointer(&f))

	sign = binary >> 31 & 0x1
	exponent = (binary >> 23) & 0xff
	fraction = binary & 0x7fffff

	signStr := "+"
	if sign == 1 {
		signStr = "-"
	}
	signInt := 1
	if sign == 1 {
		signInt = -1
	}

	fractionFloat64 := 1.0
	// fmt.Println(reflect.TypeOf(fractionFloat64))
	for i := 22; i >= 0; i-- {
		if fraction>>i&0x01 == 1 {
			fractionFloat64 += math.Pow(2, float64(i-23))
		}
	}
	/* // 別のやり方
		weight := 0.5
		for i := 22; i >= 0; i -- {
			if fraction >> i & 1 == 1 {
				fractionFloat32 += weight
			}
			weight /= 2
		}
	*/

	fmt.Printf("binary   : %01b %08b %023b\n", sign, exponent, fraction)
	fmt.Printf("sign     : %01b                                (%s)\n", sign, signStr)
	fmt.Printf("exponent : %08b                         (%d)\n" , exponent, exponent)
	fmt.Printf("fraction : 1.%023b        (%f)\n", fraction, fractionFloat64)
	fmt.Printf("result   : %f = sign * fraction * 2 ^ (exponent - 127)\n", f)
	fmt.Printf("         : %f = %b * %f * 2 ^ (%d - 127)\n", f, signInt, fractionFloat64, exponent)
	return
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: IEEE754 <floating-point-number>")
	}

	f, err := strconv.ParseFloat(os.Args[1], 32)

	if err != nil {
		fmt.Println("Command line parameter must be floating point number")
		return
	}

	// Float32Internal文字列を返す
	fmt.Print(Float32Internal(float32(f)))
}
