package main

import "testing"

func BenchmarkOptimizedUniqueSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {

		OptimizedUniqueSubstrings(moby)
	}

}

func BenchmarkUniqueSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {

		UniqueSubstrings(moby)
	}
}

var moby = `
The Project Gutenberg EBook of Moby Dick; or The Whale, by Herman
Melville

This eBook is for the use of anyone anywhere at no cost and with almost
no restrictions whatsoever.  You may copy it, give it away or re-use it
under the terms of the Project Gutenberg License included with this
eBook or online at www.gutenberg.org

`

/*
go test -run none -bench . -benchtime 3s -benchmem -v -cpuprofile p.out

BenchmarkOptimizedUniqueSubstrings-8   	     649	   5532926 ns/op	 9851415 B/op	    1309 allocs/op
BenchmarkUniqueSubstrings
BenchmarkUniqueSubstrings-8            	       6	 530396014 ns/op	 4366234 B/op	      26 allocs/op


The number of iterations that the benchmark ran. The testing package will increase this number until it believes a sufficient number of iterations have occurred.
ns/op
This stands for "nanoseconds per operation". It is the number of nanoseconds it takes to perform one operation. If your function takes in parameters, one operation is a call to your function with the parameters. The lower this number, the faster your operation function runs.
B/op
This stands for "bytes per operation". This is the amount of memory the average operation function uses. The lower the number, the less memory your function uses.
allocs/op
This stands for "allocations per operation". It signifies how many distinct memory allocations occur per operation. The less memory your function allocates the better, because memory allocation takes time.
So, in your results:
BenchmarkOptimizedUniqueSubstrings-8 was run 649 times, with each operation taking roughly 5532926 nanoseconds (or about 5.53 milliseconds), using around 9851415 bytes (or about 9.85 megabytes) per operation, and made 1309 memory allocations per operation.
BenchmarkUniqueSubstrings-8 was run 6 times, with each operation taking roughly 530396014 nanoseconds (or about 530.4 milliseconds), using around 4366234 bytes (or about 4.36 megabytes) per operation, and made 26 memory allocations per operation.
From this, you can see that BenchmarkOptimizedUniqueSubstrings-8 completes faster but uses more memory and makes more memory allocations. BenchmarkUniqueSubstrings-8 uses less memory and makes fewer memory allocations, but is slower. The "right" tradeoff depends on your specific needs and constraints
*/
