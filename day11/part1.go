package main

import (
	"log"
	"fmt"
)

func flashAt(i, j int, octopuses [][]int) (flashes int) {
	if i < 0 || i >= len(octopuses) || j < 0 || j >= len(octopuses[i]) || octopuses[i][j] == 0 {
		return 0
	}
	octopuses[i][j] += 1
	if octopuses[i][j] > 9 {
		octopuses[i][j] = 0
		flashes++
		flashes += flashAt(i-1, j-1, octopuses)
		flashes += flashAt(i-1, j, octopuses)
		flashes += flashAt(i-1, j+1, octopuses)
		flashes += flashAt(i, j-1, octopuses)
		flashes += flashAt(i, j+1, octopuses)
		flashes += flashAt(i+1, j-1, octopuses)
		flashes += flashAt(i+1, j, octopuses)
		flashes += flashAt(i+1, j+1, octopuses)
	}

	return
}

func propagateFlashes(octopuses [][]int) (flashes int) {
	for i, _ := range octopuses {
		for j, _ := range octopuses {
			if octopuses[i][j] > 9 {
				flashes += flashAt(i, j, octopuses)
			}
		}
	}

	return flashes
}

func iterateAll(octopuses [][]int) {
	for i, _ := range octopuses {
		for j, _ := range octopuses[i] {
			octopuses[i][j] += 1
		}
	}
}

func printNumFlashes(octopuses [][]int) {
	cycles := 100
	flashes := 0

	for cycle := 0; cycle < cycles; cycle++ {
		iterateAll(octopuses)
		flashes += propagateFlashes(octopuses)
	}

	fmt.Printf("There were %v flashes after %v cycles.\n", flashes, cycles)
}

func main() {
	octopuses, err := ReadOctopuses("octopuses.txt")
	if err != nil {
		log.Fatal(err)
	}

	printNumFlashes(octopuses)
}
