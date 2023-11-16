package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var value1 int
	var value2 int

	fmt.Println("Enter one number: ")
	_, err := fmt.Scanln(&value1)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	fmt.Println("Enter another number: ")
	_, err = fmt.Scanln(&value2)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	if value1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and

	if value1 == 1 && value2 == 2 {
		fmt.Println("It's true")
	} else {
		fmt.Println("It's false")
	}

	//With or

	if value1 == 0 || value2 == 2 {
		fmt.Println("It's true")
	}

	//Convert text to number

	value, err := strconv.Atoi("65")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
