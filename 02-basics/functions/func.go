package main

import "fmt"

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divideAndRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	q, r := divideAndRemainder(10, 3)
	fmt.Printf("10 divided by 3: Quotient = %d, Remainder = %d\n", q, r)
}
