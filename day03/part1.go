package main

import (
	"fmt"
	"log"
)

func printPower(diagnostics []string) {
	zeroes, ones := make([]int, len(diagnostics[0])), make([]int, len(diagnostics[0]))

	for _, d := range diagnostics {
		for i := 0; i < len(d); i++ {
			if d[i] == '0' {
				zeroes[i] += 1
			} else {
				ones[i] += 1
			}
		}
	}

	gamma, epsilon := 0, 0

	for i := 0; i < len(zeroes); i++ {
		gamma = gamma * 2
		epsilon = epsilon * 2
		if zeroes[i] > ones[i] {
			epsilon++
		} else {
			gamma++
		}
	}

	fmt.Printf("Gamma is %v and epsilon is %v.\n", gamma, epsilon)
	fmt.Printf("The power consumption is %v.\n", gamma*epsilon)
}

func main() {
	diagnostics, err := ReadDiagnostics("diagnostics.txt")
	if err != nil {
		log.Fatal(err)
	}

	printPower(diagnostics)
}
