package main

import (
	"fmt"
	"log"
	fr "Day22/FileReader"
	"Day22/Cuboid"
)

const upperBound = 50
const lowerBound = -50

func main() {
	instructions, err := fr.ReadInstructions("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := countLights(instructions)

	fmt.Printf("There are %v cubes turned on.\n", sum)
}

func countLights(instructions []*cuboid.Instruction) int {
	size := upperBound - lowerBound + 1
	lights := make([][][]bool, size)
	for i := 0; i < size; i++ {
		lights[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			lights[i][j] = make([]bool, size)
		}
	}

	for _, s := range instructions {
		for i := max(s.Cuboid.X1, lowerBound); i <= min(s.Cuboid.X2, upperBound); i++ {
			for j := max(s.Cuboid.Y1, lowerBound); j <= min(s.Cuboid.Y2, upperBound); j++ {
				for k := max(s.Cuboid.Z1, lowerBound); k <= min(s.Cuboid.Z2, upperBound); k++ {
					lights[i-lowerBound][j-lowerBound][k-lowerBound] = s.On
				}
			}
		}
	}

	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				if lights[i][j][k] {
					sum++
				}
			}
		}
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
