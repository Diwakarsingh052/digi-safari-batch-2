package main

import "fmt"

// Queue struct represents a basic queue
type Queue struct {
	items []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Println("enqueued: ", item)
}

// Dequeue removes and returns the element from the front of the queue
func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	item := queue.Dequeue()
	fmt.Println("dequeued: ", item)
}
