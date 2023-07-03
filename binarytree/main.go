// FE PMB No.9
package main

type tree [][]int

func (t tree) order(n int) {
	n -= 1
	if len(t[n]) == 2 {
		t.order(t[n][0])
		print(n+1, ", ")
		t.order(t[n][1])
	} else if len(t[n]) == 1 {
		t.order(t[n][0])
		print(n+1, ", ")
	} else {
		print(n+1, ", ")
	}
	return
}

func main() {
	t := tree{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}, {12, 13}, {14}, {}, {}, {}, {}, {}, {}, {}}
	t.order(1)
}
