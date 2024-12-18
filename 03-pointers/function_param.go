package main

import "fmt"

func doubleValue(x int) {
	x *= 2
}

func doubleValueWithPointer(x *int) {
	*x *= 2
}

func main() {
	num := 5

	doubleValue(num)
	fmt.Printf("After doubleValue: %d\n", num)

	doubleValueWithPointer(&num)
	fmt.Printf("After doubleValueWithPointer: %d\n", num)
}
