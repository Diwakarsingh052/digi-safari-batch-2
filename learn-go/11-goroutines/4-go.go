package main

import "sync"

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a ***receiver is ready***.
//send will block until there is no recv

func main() {
	wg := new(sync.WaitGroup)
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 10 // block until there is a recv
	}()

	wg.Wait()
}
