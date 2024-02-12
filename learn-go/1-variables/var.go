package main

import (
	"fmt"
)

// go mod init moduleName // initialize a go module
//go run filename.go

func main() {
	var s string // default value of string ""
	var i int = 10

	var name = "bob" // go is a statically compiled language
	//name = 100 not allowed

	fmt.Println(s, i, name)
	fmt.Printf("%q", s)

	//put related things together or all the options in this block
	var (
		//camelCase
		uName  string
		uAge   = 15
		uMarks float64
	)

	fmt.Println(uName, uAge, uMarks)
	// peek into it for design pattern
	//
	//time.Second
	//os.O_APPEND
	//http.StatusInternalServerError

}
