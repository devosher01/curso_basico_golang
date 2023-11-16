package main

import "fmt"

func equality(v1, v2 bool) bool {
	return v1 == v2
}

func different(v1, v2 bool) bool {
	return v1 != v2
}

func comparison(i1, i2 int) (ans1, ans2, ans3, ans4 bool) {
	return i1 > i2, i1 < i2, i1 >= i2, i1 <= i2
}

func logicGates(v1, v2 bool) (ans1, ans2, ans3 bool) {
	return v1 && v2, v1 || v2, !v1
}

func main() {

	var value1 bool = false
	var value2 bool = true
	var integer1 int = 5
	var integer2 int = -5

	eq := equality(value1, value2)
	ds := different(value1, value2)

	comp1, comp2, comp3, comp4 := comparison(integer1, integer2)
	andGate, orGate, notGate := logicGates(value1, value2)

	fmt.Printf("%t == %t : %t\n", value1, value2, eq)
	fmt.Printf("%t != %t : %t\n", value1, value2, ds)
	fmt.Printf("%d > %d : %t\n", integer1, integer2, comp1)
	fmt.Printf("%d < %d : %t\n", integer1, integer2, comp2)
	fmt.Printf("%d >= %d : %t\n", integer1, integer2, comp3)
	fmt.Printf("%d <= %d : %t\n", integer1, integer2, comp4)
	fmt.Printf("%t && %t : %t\n", value1, value2, andGate)
	fmt.Printf("%t || %t : %t\n", value1, value2, orGate)
	fmt.Printf("!%t : %t\n", value1, notGate)

}
