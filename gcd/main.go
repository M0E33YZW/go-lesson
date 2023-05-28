// FE PMB No.4
package main

import "fmt"

func gcd(num1 int, num2 int) int {
	x := num1
	y := num2

	for x != y {
		if x > y {
			x -= y
		} else {
			y -= x
		}
	}
	return x
}

func main() {
	fmt.Println(gcd(300, 600))
	fmt.Println(gcd(40, 8))
}
