package main

import "fmt"

func main() {
	// pointer => address of some data in the memory

	// data
	var x int
	x = 100

	// address of data
	// xPtr := &x
	var xPtr *int // the type of xPtr is "pointer to an integer"
	xPtr = &x
	fmt.Println(xPtr)

	// accessing the data in an address(pointer) - "dereferencing"
	data := *xPtr
	fmt.Println(data)

	// in other words
	fmt.Println("x == *(&x):", x == *(&x))
}
