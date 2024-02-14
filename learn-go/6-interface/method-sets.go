package main

import (
	"fmt"
)

/*
                          +-------------------------------------+
                          |              Method Set             |
                          +------------------+------------------+
                          |    Value Type    |   Pointer Type   |
+-------------------------+------------------+------------------+
| Function w/ Value Rec.  |        Yes       |       Yes        |
| ( func (t T) )          |                  |                  |
+-------------------------+------------------+------------------+
| Function w/ Ptr Rec.    |        No        |       Yes        |
| ( func (t *T) )         |                  |                  |
+-------------------------+------------------+------------------+
If a function is implemented with a value receiver (func (t T)), it can be called through a value or a pointer.
If a function is implemented with a pointer receiver (func (t *T)), it can only be called through a pointer.

var p *student
var p1 student
i = p
func (s *student) print(){


*/

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	p.name = "abc"
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func main() {

	var d1 Describer
	p1 := Person{"Sam", 25}
	p1.Describe()
	d1 = &p1
	d1.Describe()
	fmt.Println(p1)
}
