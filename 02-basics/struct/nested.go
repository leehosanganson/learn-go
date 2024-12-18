package main

import "fmt"

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Name    string
	ID      int
	Address Address
}

func main() {
	emp := Employee{
		Name: "John Doe",
		ID:   12345,
		Address: Address{
			Street:  "123 Main St",
			City:    "Reading",
			Country: "UK",
		},
	}

	fmt.Printf("%s (ID: %d) lives at %s, %s, %s\n",
		emp.Name, emp.ID, emp.Address.Street, emp.Address.City, emp.Address.Country)
}
