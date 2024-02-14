package main

import (
	"fmt"
)

// Stack represents a simple stack data structure
type Stack struct {
	items []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() int {
	// Check if the stack is empty
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	// Calculate the index of the top element in the stack
	index := len(s.items) - 1
	// Retrieve the value of the top element
	item := s.items[index]
	// Remove the top element by slicing the stack
	s.items = s.items[:index]
	// Return the value of the removed top element
	return item
}

// IsEmpty returns true if the stack is empty, otherwise false
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	// Create a new stack
	stack := Stack{}

	// Push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop and print elements
	fmt.Println("Popped:", stack.Pop())
	fmt.Println("Popped:", stack.Pop())

	// Print the size of the stack
	fmt.Println("Size of stack:", stack.Size())

	// Check if the stack is empty
	fmt.Println("Is the stack empty?", stack.IsEmpty())

}
