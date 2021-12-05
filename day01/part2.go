package main

import (
	"fmt"
	"math"
)

func main() {
	depths, err := ReadDepths("depths.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numIncreases := 0
	prevSum := math.MaxInt

	for i := 2; i < len(depths); i++ {
		newSum := depths[i] + depths[i-1] + depths[i-2]
		if newSum > prevSum {
			numIncreases++
		}
		prevSum = newSum
	}

	fmt.Printf("There are %v increases.\n", numIncreases)
}
