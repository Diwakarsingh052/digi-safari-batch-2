// Package sum is a lib package
package sum

// when we make the first letter as an upper case we export the function for other packages

var Sum int

func Add(a, b int) {
	printVal()
	Sum = a + b
}
