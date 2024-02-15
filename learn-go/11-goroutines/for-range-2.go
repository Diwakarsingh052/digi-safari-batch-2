// Import necessary packages
package main

import (
	"fmt"  // This package implements formatted I/O
	"sync" // This package provides basic synchronization primitives
)

func main() {
	// Make a channel for integers
	ch := make(chan int)

	// Use a WaitGroup to wait for all goroutines to finish
	wg := &sync.WaitGroup{}

	// Separate WaitGroup for workers sending values over the channel
	wgWorker := &sync.WaitGroup{}

	// Adding to the WaitGroup counter
	wg.Add(1)
	// Starting a goroutine
	go func() {
		// Ensure this function call is made before the goroutine exits
		defer wg.Done()

		// Loop for numbers 1-5
		for i := 1; i <= 5; i++ {
			// Adding to the WaitGroup counter for workers
			wgWorker.Add(1)
			// Start a new goroutine that sends a number to the channel
			go func(i int) {
				// Ensure this function call is made before the goroutine exits
				defer wgWorker.Done()

				// Send number to the channel
				ch <- i
			}(i) // Pass the loop variable as a function argument
		}

		// Wait until all worker goroutines finish
		//we need to block our goroutine before closing the channel because we want to make sure all the work
		// is done and finished // after closing the channel we can
		wgWorker.Wait()

		// Close the channel after all numbers are sent
		close(ch)
	}()

	// Adding to the WaitGroup counter
	wg.Add(1)
	// Start a new goroutine
	go func() {
		// Ensure this function call is made before the goroutine exits
		defer wg.Done()

		// Loop that continues until the channel is closed
		//ranging until the channel is not closed
		//range would receive all the remaining values even after the channel is closed
		for v := range ch {
			// Print each received value
			fmt.Println("Received the value from the channel ch:", v)
		}
	}()

	// Wait until all goroutines finish
	wg.Wait()
}
