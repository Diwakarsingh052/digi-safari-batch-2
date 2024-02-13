package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not in the root directory")
var ErrSomething = errors.New("some msg")

func main() {
	_, err := openFile("test.txt")
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("our custom error is found in the chain, lets create the file")
			//create the file
			//retry the operation
			//if still failed then stop this func
			return
		}
		log.Println(err)
		return

	}
}
func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		// %w means error wrapping
		return nil, fmt.Errorf("%w %w", err, ErrSomething)
	}
	return f, nil
}
