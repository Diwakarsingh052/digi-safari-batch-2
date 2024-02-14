package main

import "fmt"

func main() {
	// whenever the func is returning, defer exec
	// defer maintains a stack
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//if err != nil {
	//	panic("some panic")
	//}

	fmt.Println(4)
}
