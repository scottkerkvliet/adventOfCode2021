package main

import (
	fr "Day18/fileReader"
	sn "Day18/snailNumber"
	"fmt"
	"log"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	numbers, err := fr.ReadNumbers("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	maxSum := -1
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			maxSum = max(maxSum, sn.Add(numbers[i].Copy(), numbers[j].Copy()).GetMagnitude())
			maxSum = max(maxSum, sn.Add(numbers[j].Copy(), numbers[i].Copy()).GetMagnitude())
		}
	}

	fmt.Printf("The max sum of 2 numbers is %v\n", maxSum)
}
