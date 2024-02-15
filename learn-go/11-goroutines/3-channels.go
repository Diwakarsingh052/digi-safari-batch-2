package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
//send will block until there is no recv

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // unbuffered channel, because no size is specified
	wg.Add(2)
	go func(c chan int) {
		defer wg.Done()
		x := <-c // recv is a blocking call until there is no sender
		time.Sleep(time.Millisecond)
		fmt.Println(x)
	}(ch)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		ch <- 2 // send signal to the channel ch
	}()
	wg.Wait()

}
