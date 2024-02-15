package main

import (
	"fmt"
	"sync"
	"time"
)

// create a function work, func work(id int,wg *sync.Waitgroup){ }
// In function work print: Doing some work,id
// run a gorutine inside the work function using go func() {}()
// inside anonymous goroutine , print : running anonymous goroutine,id

// create a main function
// in main function run a loop from 1 to 3 and call work func as a goroutine, go work(i)
//guess how many goroutines would be created in above call and add it to waitgroup counter

//make sure your program runs without any deadlock, and everything gets printed out

func main() {
	// new returns a pointer to the type with its zero values initialized
	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(i, wg)
	}
	wg.Wait() // it blocks until all the goroutines who added to waitgroup counter not finishes
	fmt.Println("end of main")
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Doing some work:", i)

	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		fmt.Println("Running anonymous goRoutine:", i)
	}(i)
}
