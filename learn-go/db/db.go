package db

import (
	"errors"
	"fmt"
)

type Conn struct {
	db string
}

//factory functions
// in case host or password is empty, return an error

func NewConn(host string, password string) *Conn {
	return &Conn{db: host + password}
}

func (c *Conn) GetUser() error {
	//change the error handling in GetUser;
	//it is an incorrect way of checking if connection is present or not
	if c.db == "" {
		return errors.New("db is not set up")
	}
	fmt.Println("getting some users")
	return nil
}
func (c *Conn) Close() {
	fmt.Println("closing the connection")
	c.db = ""
}
