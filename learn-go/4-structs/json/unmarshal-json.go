package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"name"`
	Age       int    `json:"age"`
}

var rawJson = `[{"name":"John","age":30},{"name":"Rob","age":31}]`

// create a structure that can store the json after the conversion

func main() {
	var persons []person
	err := json.Unmarshal([]byte(rawJson), &persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)
}
