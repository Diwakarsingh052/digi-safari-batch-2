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
	u     user // not embedding // a field of user type
	name  string
	empId int
	pay   float64
}

func main() {
	e := employee{
		u: user{
			name:  "Bobbb",
			email: "bob@email.com",
		},
		name:  "john",
		empId: 101,
		pay:   10000000,
	}

	e.u.updateName("Bob")
	fmt.Println(e)
}
