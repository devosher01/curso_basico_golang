package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter one number")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	if number%2 != 0 {
		fmt.Println("Es impar")
	} else {
		fmt.Println("Es par")
	}

}

/*
package main

import (
	"fmt"
	"log"
)

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	var number int

	fmt.Println("Enter one number")
	_, scanErr := fmt.Scanln(&number)
	if scanErr != nil {
		log.Fatal("Error reading integer:", scanErr)
	}

	if isEven(number) {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
}
*/
