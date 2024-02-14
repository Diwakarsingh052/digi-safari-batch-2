package mysql

import (
	"database/sql"
	"fmt"
)

type Conn struct {
	db *sql.DB
}

func NewConn(db *sql.DB) Conn {
	return Conn{db: db}
}

func (c *Conn) Create(data string) {
	fmt.Println("creating a record in the db mysql", data)
}
func (c *Conn) Update(data string) {
	fmt.Println("update a record in the db mysql", data)
}

func (c *Conn) Delete(data string) {
	fmt.Println("deleting a record in the db mysql", data)
}
