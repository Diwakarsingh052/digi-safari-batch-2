package main

import "fmt"

func main() {
	//var m map[int]int // nil
	// it is an unordered collection of key value pair
	m := make(map[int]int, 100)
	m[1] = 30
	m1 := map[string]string{
		"1": "Bob",
	}

	update(m1)
	fmt.Println(m1)
}

func update(m map[string]string) {
	m["2"] = "Roy"
}
