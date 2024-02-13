package main

import "fmt"

func main() {

	// x80
	var p *int // p is a pointer which is going to store an address not a normal value
	a := 10    // x90
	p = &a     // value = x90, add of p= x80
	fmt.Println("&a &p val of p", &a, &p, p)
	//*p = 20

	//fmt.Println(a, *p) // * is a dereferencing operator // provide value at the address
	update(p)
	//fmt.Println(a, *p) // * is a dereferencing operator // provide value at the address
}

func update(ptr *int) { // val = x90 , address of ptr = x70
	fmt.Println("&ptr val of ptr", &ptr, ptr)
	*ptr = 20
}
