package main

import (
	"fmt"
)

// https://go.dev/doc/faq#methods_on_values_or_pointers
type user struct {
	name string
}

func (u1 *user) update(name string) {
	u1.name = name
}

func main() {

	var u user = user{name: "bob"}
	u.update("John")
	fmt.Println(u)
}
