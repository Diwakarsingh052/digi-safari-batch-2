package main

import (
	"fmt"
	"io"
)

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// Bigger the interface weaker the abstraction // Rob Pike
// "Don’t design with interfaces, discover them". - Rob Pike

type Reader interface {
	Read(b []byte) (int, error)
	// we need to implement all of the methods to impl the interface
	//abc()
}

type file struct {
	fileName string
}

func (f file) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them")
	return 0, nil
}

type json struct {
	fileName string
}

func (j json) Read(b []byte) (int, error) {
	fmt.Println("reading json files and processing them")
	return 0, nil
}
func (j json) name() {

}

// DoWork will accept any type of value that implements reader interface
func DoWork(r io.Reader) {
	//b := make([]byte, 100)
	r.Read(nil)
	j, ok := r.(json) // type assertion // checking if json is inside interface or not
	if ok {
		fmt.Println("calling the name from the json")
		j.name()
	}
	//we can't call methods that are not part of the interface signature
	//r.name()

}

func main() {
	var f file
	var j json

	DoWork(f)
	DoWork(j)

}
