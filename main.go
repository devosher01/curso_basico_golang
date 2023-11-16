package main

import "fmt"

func main() {
	// Constant declarations
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Integer variable declarations
	base := 12
	var height = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println("Square area:", squareArea)

	x := 10
	y := 50

	// Addition
	result := x + y
	fmt.Println("Sum:", result)

	// Subtraction
	result = y - x
	fmt.Println("Subtraction:", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Increment
	x++
	fmt.Println("Increment:", x)

	// Decrement
	x--
	fmt.Println("Decrement:", x)

	// Challenges
	// Rectangle area
	rectBase := 20
	rectHeight := 10
	rectArea := rectBase * rectHeight
	fmt.Println("Rectangle area:", rectArea)

	// Trapezoid area
	trapBase := 20
	trapTop := 10
	trapHeight := 30
	trapArea := (trapBase + trapTop) * trapHeight / 2
	fmt.Println("Trapezoid area:", trapArea)

	//Triangle area

	var baseTriangle int16 = 20
	var heightTriangle int16 = 30
	var triangleArea = (baseTriangle * heightTriangle) / 2

	fmt.Println("Triangle Area:", triangleArea)

	fmt.Println("Hello, World!")
}
