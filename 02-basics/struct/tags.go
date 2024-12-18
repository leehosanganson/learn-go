package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"published_year"`
}

func main() {
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Year:   2015,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	// Parse JSON back to struct
	var parsedBook Book
	err = json.Unmarshal(jsonData, &parsedBook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", parsedBook)
}
