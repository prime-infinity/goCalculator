package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prompt() {
	reader := bufio.NewReader(os.Stdin)
	n, err := readString("enter fitst number: ", reader)

	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}

	nn, err2 := strconv.ParseFloat(n, 64)

	if err2 != nil {
		fmt.Println("The value must be a number")
		prompt()
	} else {
		fmt.Println(nn)
	}

}

func readString(message string, r *bufio.Reader) (string, error) {
	fmt.Print(message)
	number, err := r.ReadString('\n')
	return strings.TrimSpace(number), err
}

func main() {

	prompt()

}
