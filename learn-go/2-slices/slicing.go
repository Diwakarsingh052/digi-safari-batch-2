package main

import "fmt"

func main() {
	//slicing
	a := []int{10, 20, 30, 40, 50}
	b := a[1:3] // index:len
	b = a[:4]   // start from 0th index till the length provided
	b = a[1:]   // from the 1st index till the end
	b = a[:]    // take the whole slice
	fmt.Println(b)
}
