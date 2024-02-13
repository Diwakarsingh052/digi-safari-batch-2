package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"name"`
	Age       int    `json:"age"`
	Password  string `json:"-"`
}

func main() {

	user := []person{
		{
			FirstName: "John",
			Age:       30,
			Password:  "abc",
		},
		{
			FirstName: "Rob",
			Age:       31,
			Password:  "xyz",
		},
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
