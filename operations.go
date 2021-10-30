package main

import (
	"bufio"
	"fmt"
	"os"
)

func addOperation(n []float64) {

	r := n[0] + n[1]
	results = append(results, r)
	giveResults(r)
}

func subOperation(n []float64) {
	r := n[0] - n[1]
	results = append(results, r)
	giveResults(r)
}

func divOperation(n []float64) {
	r := n[0] / n[1]
	results = append(results, r)
	giveResults(r)
}

func mulOperation(n []float64) {
	r := n[0] * n[1]
	results = append(results, r)
	giveResults(r)
}

func giveResults(r float64) {
	fmt.Println(numbers[0], operations[0], numbers[1], " = ", results[0])

	performNext()

}

func performNext() {
	reader := bufio.NewReader(os.Stdin)
	o, err := readInput("press A and to perform another calculation, press H for history: ", reader)
	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}
	switch o {
	case "A":
		fmt.Println("you are performing another calculation")
	case "H":
		fmt.Println("you are looking for history")
	default:
		fmt.Println("enter a valid answer")
		performNext()
	}

}
