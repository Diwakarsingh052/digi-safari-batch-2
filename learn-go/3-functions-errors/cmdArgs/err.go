package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//user should pass name , age, marks
	data := os.Args[1:]

	if len(data) != 3 {
		log.Println("please pass the name ,age, marks")
		return
	}
	name, age, marks := data[0], data[1], data[2]
	greeting, err := greetUser(name, age, marks)
	if err != nil {
		log.Println("cannot greet user: ", err)
		return
	}
	fmt.Println(greeting)

}

// whenever error happens, set other values to their defualts
// error should be the last type to be returned
func greetUser(nameString, ageString, marksString string) (string, error) {
	// no error = nil
	// when error (we get some kind of error message)
	//var err error // default value of error is nil
	age, err := strconv.Atoi(ageString)
	if err != nil {
		// log.Println("bad value passed for the age", err)
		return "", err
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		// log.Println("bad value passed for the marks", err)
		return "", err
	}

	s := fmt.Sprintf("Hello, %s! You are %d years old and scored %d\n", nameString, age, marks)
	return s, nil
}
