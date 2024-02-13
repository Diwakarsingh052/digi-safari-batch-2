package main

import "fmt"

type doFunc func(a, b int) int

func main() {
	op(add(10, 20))
	op(sub(10, 20))
	x := "value "
	print(x)
}

func op(do doFunc) {
	do(10, 20)
}

func add(x, y int) doFunc {
	return func(x, y int) int {
		fmt.Println(x + y)
		return x + y
	}
}

func sub(x, y int) doFunc {
	return func(x, y int) int {
		fmt.Println(x - y)
		return x + y
	}
}
