package main

import (
	"fmt"
	"log"
)

func getMaxYPosition(velocity int) int {
	return (velocity * (velocity + 1)) / 2
}

func getMaxYVelocity(yRange [2]int) int {
	target := yRange[1]
	// check if lower range is further from 0 than upper range
	if yRange[0]+yRange[1] < 0 {
		target = yRange[0]
	}

	if target < 0 {
		return (target * -1) - 1
	}
	return target
}

func printHighestYPosition(target [2][2]int) {
	yVelocity := getMaxYVelocity(target[1])
	fmt.Printf("The maximum Y position that still hits the target is %v\n", getMaxYPosition(yVelocity))
}

func main() {
	target, err := ReadTarget("target.txt")
	if err != nil {
		log.Fatal(err)
	}

	printHighestYPosition(target)
}
