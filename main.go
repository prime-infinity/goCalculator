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

	n, err := getInput("enter second number: ")
	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}
	nn, err2 := strconv.ParseFloat(n, 64)

	if err2 != nil {
		fmt.Println("The value must be a number")
		promptFirst()
	} else {
		n2 = append(n2, nn)
		fmt.Println(n1, n2, ops)
	}

}

func promptFirst() {

	n, err := getInput("enter first number: ")

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	nn, err2 := strconv.ParseFloat(n, 64)

	if err2 != nil {
		fmt.Println("The value must be a number")
		promptFirst()
	} else {

		n1 = append(n1, nn)
		enterOperation()

		/*numbers = append(numbers, nn) //appending number to slice

		if len(numbers) == 1 {
			//this is condition for first time
			enterOperation()
		}
		if len(numbers) == 2 {
			switch operation {
			case "+":
				operations = append(operations, operation) //push this operation into array
				addOperation(numbers)
			case "-":
				operations = append(operations, operation)
				subOperation(numbers)
			case "/":
				operations = append(operations, operation)
				divOperation(numbers)
			case "*":
				operations = append(operations, operation)
				mulOperation(numbers)
			default:
				fmt.Println("massive error")
				os.Exit(3)
			}
			//fmt.Println(numbers[0], numbers[1], operation)
		}*/

	}

}

func getInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	toRetrun, err := reader.ReadString('\n')
	return strings.TrimSpace(toRetrun), err
}

func main() {

	promptFirst()

}
