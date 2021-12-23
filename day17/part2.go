package main

import (
	"fmt"
	"log"
	"math"
)

func sumToN(n int) int {
	return (n * (n + 1)) / 2
}

func getXAtVelocity(step int, xVelocity int) int {
	dist := sumToN(xVelocity)
	if float64(step) < math.Abs(float64(xVelocity)) {
		dist -= sumToN(xVelocity - step)
	}

	return dist
}

func countXVelocitiesInTarget(xTarget [2]int, steps []int) (count int) {
	usedXVelocities := map[int]bool{}
	for _, step := range steps {
		upperVelocity := getXAtVelocity(step, step)
		if upperVelocity < xTarget[1] {
			diff := xTarget[1] - upperVelocity
			upperVelocity += diff / step
		}
		for {
			_, used := usedXVelocities[upperVelocity]
			if !used {
				position := getXAtVelocity(step, upperVelocity)
				if position < xTarget[0] {
					break
				}
				if position <= xTarget[1] {
					count++
					usedXVelocities[upperVelocity] = true
				}
			}
			upperVelocity--
		}
	}

	return
}

func getStepsInYTarget(yVelocity int, yTarget [2]int) (steps []int) {
	yPosition := 0
	falling := false
	for i := 0; !falling || yPosition >= yTarget[0]; i++ {
		if yPosition >= yTarget[0] && yPosition <= yTarget[1] {
			steps = append(steps, i)
		}
		yPosition += yVelocity
		yVelocity--
		falling = yVelocity < 0
	}

	return
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

func printCountInitialVelocities(target [2][2]int) {
	numVelocities := 0
	maxYVelocity := getMaxYVelocity(target[1])
	for yVelocity := maxYVelocity; yVelocity >= (maxYVelocity*-1)-1; yVelocity-- {
		steps := getStepsInYTarget(yVelocity, target[1])
		xVelocities := countXVelocitiesInTarget(target[0], steps)
		numVelocities += xVelocities
	}

	fmt.Printf("There are %v distinct initial velocities that will hit the target area.\n", numVelocities)
}

func main() {
	target, err := ReadTarget("target.txt")
	if err != nil {
		log.Fatal(err)
	}

	printCountInitialVelocities(target)
}
