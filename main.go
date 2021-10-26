package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readString(message string, r *bufio.Reader) (string, error) {
	fmt.Print(message)
	number, err := r.ReadString('\n')
	return strings.TrimSpace(number), err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	number, _ := readString("enter a number", reader)
	fmt.Println(number)
}
