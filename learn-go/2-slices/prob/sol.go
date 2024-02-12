package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := make([]int, len(x[2:5]), 100)
	copy(b, x[2:5])
	b[0] = 999
	fmt.Println(b)
	fmt.Println(x)
	b = update(b)
}

func update(a []int) []int {
	// never ever append
	return a
}
