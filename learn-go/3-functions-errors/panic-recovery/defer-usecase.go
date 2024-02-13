package main

import (
	"fmt"
	"os"
)

func main() {

	defer func(i int) {
		fmt.Println("cleaning up thing")
	}(10) // () this means we are calling the func

	fmt.Println("open a file in the next line")

	// Attempt to open the file "text.txt", using `os.Open()`. This function returns two values: a pointer to a `File` object, and an `error`.
	f, err := os.Open("text.txt")
	if err != nil {
		// handle error
	}
	defer f.Close()
	//anonymous func

	// work with the file

}
