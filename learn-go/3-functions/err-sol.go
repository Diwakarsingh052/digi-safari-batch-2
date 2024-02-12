package main

import (
	"errors"
	"fmt"
	"log"
)

// A map named 'user' holding int as key and string as value
var user = make(map[int]string)

// create a new error variable
// with msg something went bad
var ErrSomethingWentBad = errors.New("something went bad")

// ErrRecordNotFound Custom error message for record not found error
var ErrRecordNotFound = errors.New("record not found")

// FetchRecord Function tries to retrieve a value from the 'user' map
// If record is not found, it returns an error
// If record is found, it returns the retrieved value and a nil error
func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok { // if ok == false
		// wrap the newError and  ErrRecordNotFound
		return "", fmt.Errorf("%w: %w", ErrSomethingWentBad, ErrRecordNotFound)
	}
	return name, nil

}

func main() {

	n, err := FetchRecord(10)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			fmt.Println("Your account was not created")
			log.Println(err)
			return
		}
		log.Println("Internal Server Error, something wrong with the app")
		return
	}

	// if ErrRecordNotFound was added in the chain of the errors then
	// give msg to the user, your account is not created otherwise internal server error
	fmt.Println(n)
}
