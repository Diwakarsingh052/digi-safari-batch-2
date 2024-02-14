package main

import "fmt"

func main() {

	var a any

	a = 10

	a = true

	a = 10
	a = "hello"
	//s, ok := a.(int) // type assertion // asserting if bool type exist or not in any
	//if !ok {
	//	fmt.Println("type string not present in any type", s)
	//	return
	//}
	fmt.Printf("%v", a)

	show()

}

func show(i ...any) {

	fmt.Printf("%T", i)
	var r redis
	optional(r)
}

type redis struct {
	port     string
	password string
}

func optional(s redis) {

}
