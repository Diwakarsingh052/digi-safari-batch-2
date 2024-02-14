package main

import "fmt"

type Set map[int]bool

func NewSet(elems ...int) Set {
	set := make(Set)
	for _, element := range elems {
		set[element] = true
	}
	return set
}

func (s Set) Elements() []int {
	elms := make([]int, 0, len(s))
	for num := range s { // range returns two values back // key ,value
		elms = append(elms, num)
	}
	return elms
}
func (s Set) Intersect(s2 Set) []int {
	elms := make([]int, 0, len(s))
	for num := range s { // range returns two values back // key ,value
		if _, ok := s2[num]; ok {
			elms = append(elms, num)
		}
	}
	return elms
}

func (s Set) Union(s2 Set) Set {
	union := NewSet()
	for elem := range s {
		union[elem] = true
	}
	for elem := range s2 {
		union[elem] = true
	}
	return union
}

func main() {
	set1 := NewSet(1, 2, 3, 4, 5)
	set2 := NewSet(4, 5, 6, 7, 8)
	fmt.Println("Set 1:", set1.Elements())
	fmt.Println("Set 2:", set2.Elements())
	fmt.Println(set1.Intersect(set2))
	fmt.Println(set1.Union(set2))
}
