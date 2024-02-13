package main

import "fmt"

func main() {
	op(add)
	op(sub)
	x := "value "
	print(x)
}

func op(do func(a, b int)) {
	do(10, 20)

}

func add(x, y int) {
	fmt.Println(x + y)
}
func sub(x, y int) {
	fmt.Println(x - y)
}
