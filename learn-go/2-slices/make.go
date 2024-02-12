package main

import "fmt"

func main() {
	// it pre allocates a backing array of size 10
	i := make([]int, 0, 10)
	i = append(i, 100)
	inspectSlice("i", i)

}
func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()

}
