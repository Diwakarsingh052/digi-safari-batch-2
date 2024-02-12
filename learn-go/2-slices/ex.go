package main

import "fmt"

// we are going to append 100000 elms to a slice which have a initial cap of 10
// we need to calc number of allocation append made
// each time it increase the cap, we should count
// float64(currentCap - LastCap) / float64(lastCap) *100
func main() {
	data := make([]int, 0, 1000000)
	lastCap := cap(data)
	var count int

	for r := 1; r <= 1000000; r++ {
		data = append(data, r)
		if lastCap != cap(data) {
			count++
			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)

			fmt.Printf("Add [%p] Cap[%d - %v]\n", data, cap(data), capCh)
		}
	}
	fmt.Println(count)

}
