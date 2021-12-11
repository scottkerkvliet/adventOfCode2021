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
	flashes := 0
	totalOctopuses := len(octopuses) * len(octopuses[0])
	i := 0;

	for ; flashes != totalOctopuses; i++ {
		iterateAll(octopuses)
		flashes = propagateFlashes(octopuses)
	}

	fmt.Printf("All octopuses flashed at step %v.\n", i)
}

func main() {
	octopuses, err := ReadOctopuses("octopuses.txt")
	if err != nil {
		log.Fatal(err)
	}

	printNumFlashes(octopuses)
}
