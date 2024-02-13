package main

import (
	"errors"
	"fmt"
	"log"
)

// A map named 'user' holding int as key and string as value
var user = make(map[int]string)

// ErrRecordNotFound Custom error message for record not found error
var ErrRecordNotFound = errors.New("record not found")

// FetchRecord Function tries to retrieve a value from the 'user' map
// If record is not found, it returns an error
// If record is found, it returns the retrieved value and a nil error
func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok { // if ok == false

		return "", ErrRecordNotFound
	}
	return name, nil

}

func main() {

	n, err := FetchRecord(10)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(n)
}
