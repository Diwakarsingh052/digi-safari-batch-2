package main

import (
	"fmt"
	"reflect"
)

//reflection is never clear
/*
Reflection in Go is a powerful feature that allows a program to inspect and manipulate values at runtime
in ways that aren't possible with static typing.
This can be useful in a number of scenarios such as handling arbitrary data types, creating dynamic proxies, or encoding and decoding data into different formats (JSON, XML, etc.).
In Go, reflection is achieved using the reflect package.
The reflect package defines two important types: Type and Value.
Type represents the type of a Go object and provides methods to examine its structure and other attributes.
You can obtain a Type value by calling the reflect.TypeOf() function.
Value represents the value of a Go object and provides methods to inspect and modify the value.
You can obtain a Value by calling the reflect.ValueOf() function.
*/
func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
