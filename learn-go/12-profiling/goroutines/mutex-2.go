package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

//Block Profiling: Block profiling shows where goroutines are blocked waiting on synchronization primitives (including timer channels). Specifically, it shows you the stack traces where goroutines are blocked on things like channel operations, sleeping, and acquiring a lock. This is useful for finding areas of your program that are causing it to slow down due to synchronization issues. Block profiling helps to detect any contention that prevents forward progress of goroutines.
//Mutex Profiling: This focuses specifically on lock contention in your program. It shows the places where goroutines have been blocked trying to acquire a specific mutex lock. Mutex profiling provides more specific and detailed information about lock contention than the block profiler -- it will tell you which locks are most contended, how long you hold them, etc.
//In general, you would use block profiling when you want a broad look at all the places where your goroutines are getting stuck waiting for things. Mutex profiling, on the other hand, would be used when you want to focus specifically on mutex (lock) contention.

var cab int = 1

func main() {
	//
	// The 9 seconds reported in the mutex profile for m.Lock() is not in wall-clock time (i.e., real-world time). It's the total time spent by all contending goroutines waiting to acquire the lock.
	runtime.GOMAXPROCS(8)

	// Enable mutex profiling
	runtime.SetMutexProfileFraction(1)
	defer runtime.SetMutexProfileFraction(0)
	var wg = &sync.WaitGroup{}
	var m = &sync.Mutex{}
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		wg.Add(1)
		go bookCab(name, wg, m)

	}
	wg.Wait()
	// Create and write mutex profile
	f, err := os.Create("mutex.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := pprof.Lookup("mutex").WriteTo(f, 0); err != nil {
		panic(err)
	}

}

func bookCab(name string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)
	m.Lock()
	//critical section where we are using a shared resource
	// when a goroutine acquires a lock then another go routine can't access the critical section
	//until the lock is not released
	//any read , write from other goroutines would not be allowed after lock is acquired
	defer m.Unlock()
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}
	fmt.Println()
}
