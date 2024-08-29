package main

import (
	"fmt"
	"math"
)

func operate(n float64, m float64, operator string) (float64, bool) {
	// I should return Err or nil perhaps
	// let's get away with bool for now
	// so that we can tell apart a successful operation
	// from an unsuccessful one
	switch operator {
	case "+":
		return n + m, true
	case "-":
		return n - m, true
	case "*":
		return n * m, true
	case "/":
		return n / m, true
	case "^":
		return math.Pow(n, m), true
	case "sqrt":
		if n >= 0 {
			return math.Sqrt(n), true
		}
	}
	return 0, false
}

func main() {
	var n, m float64
	var operator string

	fmt.Print("First number? ")
	fmt.Scan(&n)

	fmt.Print("Operator (+, -, *, /, ^, sqrt)? ")
	fmt.Scan(&operator)

	if operator != "sqrt" {
		fmt.Print("Second number? ")
		fmt.Scan(&m)
	}

	if result, ok := operate(n, m, operator); ok == true {
		fmt.Println("=", result)
	} else {
		fmt.Println("You did something unexpected.")
	}
}
