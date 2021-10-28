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
var operation string
var counter int = 0 //counter to count the operation

func enterOperation() {

	reader := bufio.NewReader(os.Stdin)
	o, err := readInput("enter operation ( +, - , / , * ): ", reader)

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	switch o {
	case "+":
		operation = "+"
		prompt()
		//fmt.Println("you chose addition")
	case "-":
		operation = "-"
		prompt()
		//fmt.Println("you chose subtraction")
	case "*":
		operation = "*"
		prompt()
		//fmt.Println("you chose multiplication")
	case "/":
		operation = "/"
		prompt()
		//fmt.Println("you chose division")
	default:
		fmt.Println("your operation was not valid")
		enterOperation()
	}
}

func prompt() {

	counter++ //increment counter

	var opTalk string

	//this below part can be improved later using slice, by saving them in slice and iterating
	switch counter {
	case 1:
		opTalk = "first"
	case 2:
		opTalk = "second"
	case 3:
		opTalk = "third"
	default:
		opTalk = "next"
	}

	reader := bufio.NewReader(os.Stdin)
	n, err := readInput("enter "+opTalk+" number: ", reader)

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

		if len(numbers) == 1 {
			//this is condition for first time
			enterOperation()
		} else {
			switch operation {
			case "+":
				addOperation(numbers)
			case "-":
				subOperation(numbers)
			case "/":
				divOperation(numbers)
			case "*":
				mulOperation(numbers)
			default:
				fmt.Println("massive error")
				os.Exit(3)
			}
			//fmt.Println(numbers[0], numbers[1], operation)
		}
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
