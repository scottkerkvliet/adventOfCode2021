package main

import (
	"fmt"
	"log"
	"sort"
)

func getAbsValue(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getMedian(crabs []int) int {
	sort.Slice(crabs, func(i, j int) bool { return crabs[i] < crabs[j] })
	return crabs[len(crabs)/2]
}

func printFuel(crabs []int) {
	target := getMedian(crabs)
	fuel := 0
	for _, crab := range crabs {
		fuel += getAbsValue(target - crab)
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
