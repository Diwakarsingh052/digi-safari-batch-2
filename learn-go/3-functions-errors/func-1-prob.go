package main

import "fmt"

func main() {
	op(add())
	op(sub())
	x := "value "
	print(x)
}

func op(do func(a, b int)) {
	do(10, 20)
}

func add() func(int, int) {

	return func(a, b int) {
		fmt.Println(a + b)
	}
}
func sub() func(int, int) {
	return func(a, b int) {
		fmt.Println(a - b)
	}
}
