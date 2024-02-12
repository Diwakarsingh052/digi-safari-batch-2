package main

import "fmt"

func main() {

	a := 20             // create and assign the value if the variable is not already present
	a, b := 40, "hello" //a gets updated , b is created and the values are assigned

	a = 100

	n, err := fmt.Println(a, b)
	_, _ = n, err

}
