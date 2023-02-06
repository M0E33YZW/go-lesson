package main

import (
	"fmt"
	"os"
	"strconv"
)

func Float32Internal(n float32) string {

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
