package main

import "fmt"

func main() {
	// Variables Declaration

	helloMessage := "Hello"
	WorldMessage := "World"

	// Println
	fmt.Println(helloMessage, WorldMessage)
	fmt.Println(helloMessage, WorldMessage)

	// Printf

	name := "Liven"
	description := "company"

	fmt.Printf("%s is the best %s\n", name, description)
	fmt.Printf("%v is the best %v\n", name, description)

	// Sprintf

	message := fmt.Sprintf("%s is the best %s\n", name, description)
	fmt.Println(message)

	// Concatenation
	message1 := name + " is the best " + description
	fmt.Println(message1)

	// Data Types

	fmt.Printf("name: %T\n", name)
	fmt.Printf("description: %T", description)
}
