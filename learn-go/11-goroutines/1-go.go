package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	// go keyword , creates a goroutine
	// we can run function calls concurrently

	// counter to add number of goroutines
	// that we are starting or spinning up
	wg.Add(1)
	go ping(wg)
	fmt.Println("some work going on in the main func")
	wg.Wait() //block until the counter is reset to 0
	fmt.Println("end of the main")
	//no guesses
	//time.Sleep(4 * time.Second)

}

func ping(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter
	time.Sleep(3 * time.Second)
	fmt.Println("ping")
}
