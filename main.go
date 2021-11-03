package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//the first and second set of numbers,
//operations and results are all stored in their respective slices
var n1 = []float64{}
var n2 = []float64{}
var ops = []string{} //operations
var results = []float64{}

func enterOperation() {

	o, err := getInput("enter operation ( +, - , / , * ): ")

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	switch o {
	case "+":
		ops = append(ops, "+")
	case "-":
		ops = append(ops, "-")
	case "*":
		ops = append(ops, "*")
	case "/":
		ops = append(ops, "/")
	default:
		fmt.Println("your operation was not valid")
		enterOperation()
	}
	promptSecond()
}

func promptSecond() {

	returnNumber := scrutInput("enter second number: ", 2)
	n2 = append(n2, returnNumber)

	len := len(n1) - 1

	result := performOperation(n1[len], n2[len], ops[len])

	results = append(results, result)
	giveResults()
}

func giveResults() {
	fmt.Println(results[len(results)-1])
}

func promptFirst() {

	returnNumber := scrutInput("enter first number: ", 1)

	n1 = append(n1, returnNumber)
	enterOperation()
}

func getInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	toRetrun, err := reader.ReadString('\n')
	return strings.TrimSpace(toRetrun), err
}

func scrutInput(message string, step int) float64 {
	n, err := getInput(message)

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}
	nn, err2 := strconv.ParseFloat(n, 64)
	if err2 != nil {
		fmt.Println("The value must be a number")
		if step == 1 {
			promptFirst()
		}
		if step == 2 {
			promptSecond()
		}

	}
	return nn

}

func main() {

	promptFirst()

}
