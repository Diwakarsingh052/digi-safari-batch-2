package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("sending a notification to %s %s %s", u.name, u.email, string(p))
	return len(p), nil
}

func main() {
	u := user{
		name:  "bob",
		email: "bob@email.com",
	}
	// os.Stdout instead of u
	l := log.New(u, "sales: ", log.Lshortfile)
	l.Println("this is a sample log")
}
