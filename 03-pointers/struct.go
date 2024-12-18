package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// If not using pointer receiver, the method would be called on a copy of the struct
// which means the original struct (p.Age) would not be modified
func (p *Person) haveBirthday() {
	p.Age++
}

func main() {
	person := &Person{Name: "Alice", Age: 30}

	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	person.haveBirthday()
	fmt.Printf("After birthday: %s is %d years old\n", person.Name, person.Age)
}
