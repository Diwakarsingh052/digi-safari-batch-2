package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[3:5] // 40, 50, // l=2 , c =4
	inspectSlice("b", b)
	b = append(b, 777, 888, 999) // 40, 50, 60, 777,888,999
	inspectSlice("b", b)
	inspectSlice("x", x)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()

}
