package main

import "fmt"

func authFunction(username, password string) bool {
	if username == "oscar" && password == "qwe" {
		fmt.Println("Correct Data")
		return true
	}
	println("wrong data")
	return false
}

func main() {
	var username string
	var password string
	fmt.Println("Enter your Username")
	_, err := fmt.Scanln(&username)
	if err != nil {
		println("There was a user error", err)
		return
	}

	fmt.Println("Enter your password")
	_, err2 := fmt.Scanln(&password)
	if err2 != nil {
		println("There was a password error", err)
		return
	}

	result := authFunction(username, password)

	if result {
		fmt.Println("Correctly authenticated")
	} else {
		fmt.Println("There was a error, try it again")
	}

}

/*
package main

import "fmt"

func authFunction(username, password string) bool {
	return username == "oscar" && password == "qwe"
}

func main() {
	var username string
	var password string

	fmt.Println("Enter your Username")
	_, usernameErr := fmt.Scanln(&username)
	if usernameErr != nil {
		fmt.Println("There was a user error:", usernameErr)
		return
	}

	fmt.Println("Enter your password")
	_, passwordErr := fmt.Scanln(&password)
	if passwordErr != nil {
		fmt.Println("There was a password error:", passwordErr)
		return
	}

	if authFunction(username, password) {
		fmt.Println("Correctly authenticated")
	} else {
		fmt.Println("Authentication failed, try again")
	}
}

*/
