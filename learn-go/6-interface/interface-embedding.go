package main

import "io"

type student struct {
}

func (s student) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s student) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}
func main() {
	var s student
	accept(s)
}

// any type that implements both Read and Write methods matches io.ReadWriter interface.
// If student has both Read and Write method, this assignment is valid
func accept(rw io.ReadWriter) {

}
