package main

import "fmt"

const (
	// BURGER = 0b0001
	// POTATO = 0b0010
	// COFFEE = 0b0100
	// NUGGET = 0b1000

	/* iota:
	0からの連番を振る
	BURGER = (1 << iota) / 2
	1からの連番を振る
	BURGER = 1 << iota
	*/

	BURGER = (1 << iota) >> 1 // 0b0001
	POTATO                    // 0b0010
	COFFEE                    // 0b0100
	NUGGET                    // 0b1000
)

type Order int

func (o Order) has(n int) bool {
	return int(o) & n != 0
}

func main() {
	order := Order(0b0011)

/* 	if order&BURGER != 0 {
		fmt.Println("BURGER")
	}
	if order&BURGER != 0 {
		fmt.Println("POTATO")
	}
	if order&BURGER != 0 {
		fmt.Println("COFFEE")
	}
	if order&BURGER != 0 {
		fmt.Println("NUGGET")
	} */
	if order.has(BURGER) != 0 {
		fmt.Println("BURGER")
	}
	/* 
	if order.has(POTATO) != 0 {
		fmt.Println("POTATO")
	}
	if order.has(COFFEE) != 0 {
		fmt.Println("COFFEE")
	}
	if order.has(NUGGET) != 0 {
		fmt.Println("NUGGET")
	} */
}
