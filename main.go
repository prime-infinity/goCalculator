package main

import (
	"bufio"
	"fmt"
	"os"
)

func readString(message string, r *bufio.Reader) string /*, error*/ {
	fmt.Print(message)
	return message
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	number /*, _*/ := readString("enter a number", reader)
	fmt.Println(number)
}
