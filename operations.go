package main

import "fmt"

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
	fmt.Println(numbers, operations, results)
}
