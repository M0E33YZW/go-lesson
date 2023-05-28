// FE PMB No.5
package main

import (
	"fmt"
	"math"
)

func calc(x float64, y float64) float64 {
	c := math.Pow(x, 2.0)
	d := math.Pow(y, 2.0)
	return math.Pow(c+d, 0.5)
}

func main() {
	fmt.Println(calc(float64(2), float64(3)))
}
