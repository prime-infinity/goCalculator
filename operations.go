package main

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

/*func performNext() {
	reader := bufio.NewReader(os.Stdin)
	o, err := readInput("press X to perform calculation on result, A to perform another calculation, press H for calculation history: ", reader)
	if err != nil {
		fmt.Println("there was a fatal problem")
		os.Exit(3)
	}
	switch o {
	case "X", "x":
		fmt.Println("you are performing calculation on result")
	case "A", "a":
		fmt.Println("you are performing another calculation")
		counter = 0 //make counter 0 again
		prompt()
	case "H", "h":
		fmt.Println("you are looking for history not available yet")
	default:
		fmt.Println("enter a valid answer")
		performNext()
	}

}*/
