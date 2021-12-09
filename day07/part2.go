package main

import (
	"fmt"
	"log"
)

func getAbsValue(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getFuelCost(distance int) int {
	return ((distance + 1) * distance) / 2
}

func getMean(crabs []int) int {
	var sum float64 = 0
	for _, crab := range crabs {
		sum += float64(crab)
	}

	avg := sum / float64(len(crabs))
	mean := int(avg)
	if avg == float64(mean) {
		return mean
	}

	// Determine whether we should have rounded up instead of down
	// Based on if mean is below median, not decimals from avg
	crabsFurther := 0
	for _, crab := range crabs {
		if float64(crab) > avg {
			crabsFurther++
		}
	}
	if crabsFurther > len(crabs)/2 {
		mean++
	}

	return mean
}

func printFuel(crabs []int) {
	target := getMean(crabs)
	fuel := 0
	for _, crab := range crabs {
		fuel += getFuelCost(getAbsValue(target - crab))
	}

	fmt.Printf("Crabs aligned at %v. It took %v fuel.\n", target, fuel)
}

func main() {
	crabs, err := ReadCrabs("crabs.txt")
	if err != nil {
		log.Fatal(err)
	}

	printFuel(crabs)
}
