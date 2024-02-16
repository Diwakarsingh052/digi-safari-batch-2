package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// function that locks and unlocks shared mutex in infinite loop
func lockMutexInLoop(mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		mutex.Lock()
		time.Sleep(5 * time.Millisecond)
		mutex.Unlock()
	}
}

func main() {
	runtime.SetBlockProfileRate(1)
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	// Start goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go lockMutexInLoop(mutex, wg)
	}
	wg.Wait()
	// Wait for a while to allow some block profile data to be collected

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
