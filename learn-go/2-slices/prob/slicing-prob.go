package main

import "fmt"

// a:= 10,20,30,40,50,60  // backing array
//b[0] = 999 // b is sharing the same backing array with a slice, so any updates in b will also result changes in A slice
// a ,(b)-> [10,(999,30,40,50),60] // b is not going to get its own copy

func main() {

	i := []int{10, 20, 30, 40, 50, 60, 70}
	b := i[1:5] //20, 30, 40, 50
	b[0] = 999
	fmt.Println(b)
	fmt.Println(i)

}
