package main

import (
	"errors"
	"fmt"
	"log"
)

//Error should be suffixed in the name of struct, for e.g. QueryError
// the struct should not be used as a normal struct to store some data

type QueryError struct {
	Func  string
	Input string
	Err   error
}

func (q *QueryError) Error() string {
	return "main." + q.Func + " searching " + q.Input + ":" + q.Err.Error()
}

var ErrNotFound = errors.New("not found")

func SearchSomething(s string) error {
	if s == "simple error" {
		return errors.New("a simple err, not using a struct")
	}
	if s != "value found" {
		return &QueryError{
			Func:  "SearchSomething",
			Input: s,
			Err:   errors.New("SearchSomething failed"),
		}
	}

	return nil
}
func main() {

	var q *QueryError // nil
	err := SearchSomething("simple error")
	if err != nil {
		// errors.As is going to look for a custom type added in the chain or not
		// if a custom type is present in the chain, it would assign the value to the variable
		// passed as argument to errors.As
		if errors.As(err, &q) {
			fmt.Println("custom error found in the chain", q.Func)
		}
		log.Println(err)
	}

}

//fmt.Println(strconv.Atoi("abc"))
//	fmt.Println(strconv.Atoi("xyz"))
//	fmt.Println(strconv.ParseInt("efg", 10, 64))
