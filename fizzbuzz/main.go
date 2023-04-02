package main

import "fmt"

var a int = 3
var b int = 5

func fizzBuzz(num int) {
	if num%a == 0 && num%b == 0 {
		fmt.Printf("%d: %dと%dで割り切れる\n", num, a, b)
	} else if num%a == 0 && num%b != 0 {
		fmt.Printf("%d: %dで割り切れる\n", num, a)
	} else if num%a != 0 && num%b == 0 {
		fmt.Printf("%d: %dで割り切れる\n", num, b)
	} else {
		fmt.Printf("%d: %dでも%dでも割り切れない\n", num, a, b)
	}
}

func main() {
	for i := 6; i < 20; i++ {
		fizzBuzz(i)
	}
}
