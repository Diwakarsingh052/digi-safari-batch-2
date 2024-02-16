// Import necessary packages
package main

import (
	"context" // For managing multiple goroutine lifetimes
	"fmt"     // Implements formatted I/O.
	"sync"    // Provides basic synchronization primitives
	"time"    // For time measurements and operations
)

// This function simulates a slow process that generates an integer value
func slowFunction() int {
	// Pauses the function for 5 seconds
	time.Sleep(5 * time.Second)
	return 42 // Returns the integer 42
}

func main() {
	// Create a new channel that can send and receive integers
	ch := make(chan int)

	// Create a WaitGroup. A WaitGroup waits for a collection of goroutines to finish.
	wg := new(sync.WaitGroup)

	// Create a new context. This is essentially a signal to all
	// goroutines that they should stop what they are doing and return.
	ctx := context.Background()

	// Create a new context that will automatically cancel after 1 millisecond
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)

	// Add a job to the WaitGroup
	wg.Add(1)

	// Start a goroutine (a lightweight thread)
	go func() {
		// Schedule the call to Done to tell the WaitGroup that we have finished executing
		defer wg.Done()

		// Attempt to send the result of slowFunction() on the channel if timeout is not over
		select {
		case <-ctx.Done(): // If the context indicates done (timeout reached), skip sending process
			fmt.Println("The timer expired before the process finished")
			return

		case ch <- slowFunction(): // If timeout not hit yet, run slowFunction() and send result to the channel
			fmt.Println("Process finished in time and sent the result")

		}
		// Print that the sender goroutine has finished
		fmt.Println("Sender goroutine finished ")
	}()

	// Call the cancellation function to release resources associated with it
	defer cancel()

	// Wait some time for the result coming from the channel or until the timeout hits
	select {
	case res := <-ch: // If the result arrived in time
		fmt.Println("Received value from slow function:", res)
	case <-ctx.Done(): // If the context finished before the result arrived (timeout)
		fmt.Println("Operation timed out!")
	}

	// Print that there are no more values to receive
	fmt.Println("no more receivers")

	// Wait until all jobs added to WaitGroup are done
	wg.Wait()
}
