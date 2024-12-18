package main

import "fmt"

func main() {
	// Array
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fmt.Println("Fruits:", fruits)

	// Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Append to slice
	numbers = append(numbers, 6, 7)
	fmt.Println("Updated numbers:", numbers)

	// Slice of slice
	subset := numbers[2:5]
	fmt.Println("Subset:", subset)
}
