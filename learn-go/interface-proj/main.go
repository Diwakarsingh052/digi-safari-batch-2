package main

import (
	"interface-proj/data"
	"interface-proj/data/mysql"
	"interface-proj/data/postgres"
)

func main() {

	m := mysql.NewConn(nil)
	p := postgres.NewConn(nil)
	s := data.NewStore(&m)
	s.Create("random data")

	s = data.NewStore(&p)
	s.Delete("fake data")

}
