package main

import "fmt"

// When you're working with a pointer to a struct (like l in your previous example),
// what's stored in memory is the memory address where the struct's values start
// (base address).
// When you create an instance of this struct (s := MyStruct{}),
// Go will allocate a contiguous block of memory with enough space to hold all these fields.
type list struct {
	num []int
}

func (l1 *list) addToList() {
	l1.num = append(l1.num, 40, 50)
}

func main() {

	l := &list{num: []int{10, 20, 30}} // l = 3, c=3

	l.addToList()

	fmt.Println(l)
}
