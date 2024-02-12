package main

import "fmt"

type student struct {
	name  string // fields
	age   int
	marks []int
}

func (s student) showValues() { // func (recv)Name(args)returnTypes
	fmt.Println(s.name, s.age, s.marks)
}

func main() {
	var s1 student
	s1.name = "some Name"
	s1.marks = []int{10, 20, 30, 50}
	for i, v := range s1.marks {
		fmt.Println(i, v)
	}

	// s1 value would be copied to the receiver of the method
	s1.showValues()
}
