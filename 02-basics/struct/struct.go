package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a struct instance
	p1 := Person{Name: "Alice", Age: 30}

	// Access struct fields
	fmt.Printf("%s is %d years old\n", p1.Name, p1.Age)

	// Create a struct with zero values
	var p2 Person
	p2.Name = "Bob"
	p2.Age = 25
	fmt.Printf("%s is %d years old\n", p2.Name, p2.Age)
}
