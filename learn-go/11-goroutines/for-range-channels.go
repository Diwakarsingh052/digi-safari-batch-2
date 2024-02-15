package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// it ranges over the channel to recv the values, but it is going to stop when we close the channel
		// close the channel when work is done
		for v := range ch {
			fmt.Println("recv the value from the channel ch", v)
		}
	}()

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // this would close the channel // it would signal no more values would be sent
	// range can read the remaining values and quit
	wg.Wait()
}
