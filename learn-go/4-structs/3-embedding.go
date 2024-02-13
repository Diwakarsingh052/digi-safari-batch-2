package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateName(name string) {
	fmt.Println("user")
	u.name = name
}

type employee struct {
	user  // anonymous field // embedding
	name  string
	empId int
	pay   float64
}

func (e *employee) updateName(name string) {
	fmt.Println("emp")
	e.name = name
}

func main() {
	e := employee{
		user: user{
			name:  "Bobbb",
			email: "bob@email.com",
		},
		name:  "john",
		empId: 101,
		pay:   10000000,
	}
	var u user
	e.user = u

	e.updateName("Bob")
	fmt.Println(e)
}
