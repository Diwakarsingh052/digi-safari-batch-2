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
	greetUser(name, age, marks)

}

func greetUser(name, ageString, marksString string) {
	// no error = nil
	// when error (we get some kind of error message)
	//var err error // default value of error is nil
	// if a func returns an err , handle the error first before doing anything
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("bad value passed for the age", err)
		return
	}
	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println("bad value passed for the marks", err)
		return
	}
	fmt.Println(name, age, marks)

	// from this function report error back to main if it happened

}
