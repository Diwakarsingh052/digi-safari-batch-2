package main

import "fmt"

type order struct {
	orderId int
}

func main() {
	o := []order{
		{
			orderId: 1,
		},
		{
			orderId: 2,
		},
	} // x80
	o = updateStructOrder(o)
	fmt.Println(o)
}

func updateStructOrder(o1 []order) []order {
	o1[0] = order{orderId: 100}
	o1 = append(o1, order{orderId: 3}) // x90
	return o1
}
