package main

import "fmt"

func greeting(name string) { // name is the parameter
	fmt.Println("Hello", name)
}

func tripleArgument(a, b int, c string) { // a, b, c are the arguments
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // int is the return type
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	greeting("Oscar")             // Hello Oscar
	greeting("Fernando")          // Hello Fernando
	tripleArgument(2, 3, "Hello") // 2 3 Hello

	value := returnValue(2)      // value = 4
	fmt.Println("Value:", value) // Value: 4

	value1, value2 := doubleReturn(5)
	fmt.Println("Value 1:", value1, " ", "Value 2:", value2)
}
