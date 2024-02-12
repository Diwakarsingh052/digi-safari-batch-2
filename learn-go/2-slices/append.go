package main

import "fmt"

func main() {
	//If the capacity of s is not large enough to fit the additional values,
	//append allocates a new, sufficiently large underlying array that fits
	//both the existing slice elements and the additional values.
	//Otherwise, append re-uses the underlying array.
	// [*p,len,capacity]
	i := []int{10, 20, 30, 40} // l = 4 , c = 4
	fmt.Println(len(i), cap(i))
	i = append(i, 50, 60) // l = 6 , c = 8
	fmt.Println(len(i), cap(i))
	i = append(i, 70) // l = 7 , c = 8
	fmt.Println(len(i), cap(i))
	fmt.Println(i)
}
