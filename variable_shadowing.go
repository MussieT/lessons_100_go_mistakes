package main

import "fmt"

// This use case usually happens on the error variable

// Wrong code
func ShadowedVariable() {
	var result int
	operation := "addition"

	if operation == "addition" {
		a := 10
		b := 11

		result := a + b

		fmt.Println("result: ", result)
	} else {
		a := 10
		b := 11

		result := a - b

		fmt.Println("result: ", result)

	}

	fmt.Println("result inside shadowed: ", result)
}

// Valid code
func UnShadowedVariable() {
	var result, a, b int
	operation := "addition"

	if operation == "addition" {
		a = 10
		b = 11

		result = a + b

		fmt.Println("result: ", result)
	} else {
		a = 10
		b = 11

		result = a - b

		fmt.Println("result: ", result)

	}

	fmt.Println("result inside not shadowed: ", result)
}
