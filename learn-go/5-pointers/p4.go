package main

import "fmt"

// append allocates new memory in this function and that new memory is not returned back to the list
func addToList(s []int) {
	fmt.Printf("s %p\n", s)
	s = append(s, 888, 999) // updated the copy
	fmt.Printf("after append s %p\n", s)

}
func main() {
	list := []int{10, 20} // l = 2 , c =2
	fmt.Printf("list %p\n", list)
	// Call addToList, expecting to add 888 and 999 to the end of 'list'
	// However, the function doesn't actually modify 'list', so it remains as [10, 20]
	addToList(list)
	fmt.Printf("after append list %p\n", list)
	fmt.Println(list)
}
