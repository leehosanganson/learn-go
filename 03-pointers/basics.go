package main

import "fmt"

func main() {
	// Declare a variable
	x := 10

	// Declare a pointer to an int
	var p *int

	// Assign the address of x to p
	p = &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p: %p\n", p)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// Modify the value through the pointer
	*p = 20
	fmt.Printf("New value of x: %d\n", x)
}
