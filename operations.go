package main

import (
	"fmt"
)

func performOperation(n1 float64, n2 float64, ops string) float64 {

	switch ops {
	case "+":
		return n1 + n2
	case "/":
		return n1 / n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	}
	return 0
}

func performNext() {
	o, _ := getInput("press X to perform calculation on result, A to perform another calculation, press H for calculation history: ")
	switch o {
	case "X", "x":
		fmt.Println("you are performing calculation on result")
		n1 = append(n1, results[len(results)-1])
		enterOperation()
	case "A", "a":
		fmt.Println("you are performing another calculation")
	case "H", "h":
		fmt.Println("you are looking for history not available yet")
	default:
		fmt.Println("enter a valid answer")
		performNext()
	}

}
