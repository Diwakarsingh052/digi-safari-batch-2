package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var wgBook = sync.WaitGroup{}

// Theater represents a theater with a specific number of seats
type Theater struct {
	Seats   int          // Total seats available in the theater
	invoice chan string  // Channel to store the name of the person whose seat is booked
	rw      sync.RWMutex // Read-write Mutex for synchronizing shared access to Seats
}

// checkSeats method checks the available Seats in the Theater
func (t *Theater) checkSeats() {
	defer wg.Done() // Decrement the counter when checkSeats is done executing

	// Acquire a lock for reading
	t.rw.RLock()
	// Releases the read lock when func completes
	defer t.rw.RUnlock()

	fmt.Println("Available Seats:", t.Seats)
}

// bookSeat method attempts to book a seat in the Theater
func (t *Theater) bookSeat(name string) {
	defer wgBook.Done() // Decrement the counter when bookSeat is done executing

	// Acquire a write lock
	t.rw.Lock()
	// Releases the write lock when func completes
	defer t.rw.Unlock()

	// If there are seats available
	if t.Seats > 0 {
		// Simulate a seat booking-making process
		fmt.Println("Seat is available for", name)
		time.Sleep(2 * time.Second)
		fmt.Println("Booking confirmed", name)

		t.Seats--         // Decrement available seats
		t.invoice <- name // Send a person's name to the invoice channel
	} else {
		fmt.Println("No seats available for", name) // Inform that no seats are available
	}
}

// Main routine
func main() {
	// Create a new Theater
	t := Theater{
		Seats:   4,                 // With 1 seat
		invoice: make(chan string), // Create the invoice channel
	}
	done := make(chan struct{}) // Create a done channel for handling completion of seat bookings

	// Start 2 checkSeat routines
	for i := 1; i <= 6; i++ {
		wg.Add(1) // Increment wait group counter
		go t.checkSeats()
	}

	// Start 2 bookSeat routines
	for i := 1; i <= 6; i++ {
		wgBook.Add(1) // Increment wait group counter
		go t.bookSeat("User " + strconv.Itoa(i))
	}

	// When all bookings are completed, it signals the 'done' channel
	go func() {
		wgBook.Wait()
		close(done)
	}()

	wg.Add(1)
	go t.printInvoice(done)

	// Checks Seats availability after booking operation has completed
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go t.checkSeats()
	}

	wg.Wait() // Wait for all active go routines to complete their tasks
}

// printInvoice method prints the invoice for all booked seats
func (t *Theater) printInvoice(done chan struct{}) {
	defer wg.Done() // Decrement the counter when func is done executing
	for {
		select {
		case name := <-t.invoice:
			// Print Invoice for the person whose name was sent through the invoice channel
			fmt.Printf("Invoice is sent to %s\n", name)
		case <-done:
			// If all invoices are sent, stops the loop and returns from the function.
			fmt.Println("All invoices are sent")
			return
		}
	}
}
