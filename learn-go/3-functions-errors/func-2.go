package main

import "fmt"

func main() {
	r := op(add, 10, 20)
	fmt.Println(r)

	r = op(sub, 100, 20)
	fmt.Println(r)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func op(do func(a, b int) int, a, b int) int {
	return do(a, b)
}
