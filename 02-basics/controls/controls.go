package main

import "fmt"

func main() {
	// If statement
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Switch statement
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
