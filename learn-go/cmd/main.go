package main

import (
	"learn-go/db"
	"log"
)

func main() {
	//sum.Add(10, 20)
	//sum.Sum = 100
	//fmt.Println(sum.Sum)

	c := db.NewConn("host ", "password")
	defer c.Close()

	err := c.GetUser()
	if err != nil {
		log.Println(err)
		return
	}

}
