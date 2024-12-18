package main

import "fmt"

func main() {
	// Create and initialize a map
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	// Add a new key-value pair
	ages["David"] = 28

	// Access and print values
	fmt.Println("Bob's age:", ages["Bob"])

	// Check if a key exists
	if age, exists := ages["Eve"]; exists {
		fmt.Println("Eve's age:", age)
	} else {
		fmt.Println("Eve not found in the map")
	}

	// Iterate over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
