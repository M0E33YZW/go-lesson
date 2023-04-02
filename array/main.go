package main

import "fmt"

func makeNewArray(in []int) {
	out := []int{}
	out = append ([] int{in[0]})

	for i := 1; i < len(in); i++ {
		tail := out[len(out) - 1]
		out = append(out, tail + in[i])
	}
	fmt.Print(out)
}

func main() {
	array := []int{3, 2, 1, 6, 5, 4}

	if len(array) >= 2 {
		makeNewArray(array)
	}
}
