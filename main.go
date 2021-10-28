package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//the numbers that are to be entered, stored as slice
var numbers = []float64{}

func enterOperation() {

	reader := bufio.NewReader(os.Stdin)
	o, err := readInput("enter operation ( +, - , / , * ): ", reader)

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	switch o {
	case "+":
		fmt.Println("you chose addition")
	case "-":
		fmt.Println("you chose subtraction")
	case "*":
		fmt.Println("you chose multiplication")
	case "/":
		fmt.Println("you chose division")
	default:
		fmt.Println("your operation was not valid")
		enterOperation()
	}
}

func prompt() {
	reader := bufio.NewReader(os.Stdin)
	n, err := readInput("enter fitst number: ", reader)

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	nn, err2 := strconv.ParseFloat(n, 64)

	if err2 != nil {
		fmt.Println("The value must be a number")
		prompt()
	} else {
		numbers = append(numbers, nn) //appending number to slice
		fmt.Println(numbers)
		enterOperation()
	}

}

func readInput(message string, r *bufio.Reader) (string, error) {
	fmt.Print(message)
	toRetrun, err := r.ReadString('\n')
	return strings.TrimSpace(toRetrun), err
}

func main() {

	prompt()
	//fmt.Printf("%T,%v", numbers, numbers)

}
