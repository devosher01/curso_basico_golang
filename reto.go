package main

import "fmt"

func findSquareAndRectangleArea(base, height int) int {
	return base * height
}

func findTriangleArea(base, height int) float64 {
	return float64(base) * float64(height) / 2
}

func findTrapezoidArea(base, base2, height int) float64 {
	return float64(base+base2) * float64(height) / 2
}

func main() {

	// Square, Rectangle and Triangle
	baseC := 3
	height := 5

	squareAndRectangleArea := findSquareAndRectangleArea(baseC, height)
	fmt.Println("squareAndRectangleArea:", squareAndRectangleArea)

	TriangleArea := findTriangleArea(baseC, height)
	fmt.Println("triangleArea:", TriangleArea)

	// Trapezoid
	baseT := 3
	base2 := 5
	heightT := 5

	TrapezoidArea := findTrapezoidArea(baseT, base2, heightT)
	fmt.Println("trapezoidArea:", TrapezoidArea)

}
