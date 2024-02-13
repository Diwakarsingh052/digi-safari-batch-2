package main

import "fmt"

func slice(a []int) {
	a[2] = 10
	fmt.Println("slice finished")
}

func doSomeStuff() {
	defer recoveryFunc()
	slice([]int{10})
	fmt.Println("doSomeStuff finished")
}

func main() {

	doSomeStuff()
	fmt.Println("end of the main")
}

// `recoveryFunc` uses `recover` function which regains control of a panicking goroutine, effectively
// preventing the program from crashing.
// Recover is only useful inside deferred functions.
// If the goroutine is panicking, recover will capture the panic value and stop the panic.
// If the goroutine is not panicking, recover does nothing.
func recoveryFunc() {
	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.
	if r := recover(); r != nil {
		// If `recover` captured a panic, it returns the panic value.
		// Here we print it.
		fmt.Println("Recovered from panic:", r)
	}

}
