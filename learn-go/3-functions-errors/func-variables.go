package main

import "fmt"

type greet func(msg string)

func main() {
	f := func(msg string) {
		fmt.Println("some stuff", msg)
	}
	f("this is a test func")
	printSomething(f)
}

func printSomething(f1 greet) {

}
