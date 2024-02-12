package main

import "fmt"

type money int64 // newType is money, underlying type is int

func main() {
	var i money = 1000
	var b int64 = int64(i)
	fmt.Println(b)
	//var x rune
}
