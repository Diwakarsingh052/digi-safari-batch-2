package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	// To keep track of if the goroutine work is finished or not,
	// we need another goroutine to close the channel (done).
	// This pattern is useful in case of multiple producers sending data on
	// the channel and single consumer reading from it. Closing the channel
	// signals the consumer about the completion of data sending by all producers.
	wgWorker := sync.WaitGroup{}

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(4 * time.Second)
		c1 <- "1"
		c1 <- "4"
	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(2 * time.Second)
		c2 <- "2"
	}()
	go func() {
		defer wgWorker.Done()
		c3 <- "3"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done) // We are closing the channel (done) when all goroutines are finished sending.
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Infinite for loop with select is useful here because we don't know how many value
		// we would receive from the channels
		// whatever channel sends the data first, that case will execute.
		// This gives us concurrency and we don't have to wait for a particular channel to finish.
		// Also, it allows the easy addition of more channels in the future.
		for {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			// possible cases are chan recv , send , default
			select {
			case x := <-c1:
				fmt.Println(x)
			case x := <-c2:
				fmt.Println(x)
			case x := <-c3:
				fmt.Println(x)
			case <-done: // this case will execute when channel is closed, signalling all work done
				fmt.Println("work is finished")
				return
			}
		}
	}()
	wg.Wait()
	// Create a profile file
	f, err := os.Create("block_profile.pprof")
	if err != nil {
		log.Fatal("could not create block profile: ", err)
	}
	defer f.Close()

	// Write the current blocking profile to the file
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		log.Fatal("could not write block profile: ", err)
	}
}

//wgWorker.Wait() in the main.main.func4 function: This command is causing the goroutine to wait and blocked it for 4 seconds. The WaitGroup is used to ensure that the program should not proceed until the 3 goroutines sending on c1, c2 and c3 channels are finished. This is expected behavior, as WaitGroup.Wait() is designed to block until all tasks are done.
//select in the main.main.func5 function: The select statement is blocked and pauses for 4 seconds waiting for input on the c1, c2, c3 channels. The select statement is designed to wait until there is a communication that can be either send or receive ready to proceed.
//wg.Wait() in the main.main function: This line also blocked the execution for 4 seconds. Here, the WaitGroup wg is waiting for the other goroutines to catch up. Specifically, it's waiting for the goroutines running wgWorker.Wait() to finish and for the select statement to finish processing channels.
//In this sample, all the blocking times are expected and are part of the normal program workflow. However, if you were to observe surprisingly long blocking times or blocking times in a part of the program where you didn't expect, then that could be indicative of a performance issue where some goroutines are waiting on others too much.
