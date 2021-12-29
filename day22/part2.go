package main

import (
	"fmt"
	"log"
	fr "Day22/FileReader"
	"Day22/Cuboid"
)

func main() {
	instructions, err := fr.ReadInstructions("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}

	onCuboids := applyInstructions(instructions)
	sum := 0
	for _, c := range onCuboids {
		sum += c.GetVolume()
	}

	fmt.Printf("There are %v cubes turned on.\n", sum)
	fmt.Printf("There were %v cuboids to parse.\n", len(onCuboids))
}

func applyInstructions(instructions []*cuboid.Instruction) []*cuboid.Cuboid {
	var onCuboids []*cuboid.Cuboid
	for _, i := range instructions {
		newOnCuboids := make([]*cuboid.Cuboid, 0, len(onCuboids)+1)
		for _, c := range onCuboids {
			newCuboids := cuboid.RemoveFromCuboid(c, i.Cuboid)
			for _, nc := range newCuboids {
				newOnCuboids = append(newOnCuboids, nc)
			}
		}
		if i.On {
			newOnCuboids = append(newOnCuboids, i.Cuboid)
		}
		onCuboids = newOnCuboids
	}

	return onCuboids
}


func trimCuboid(c *cuboid.Cuboid) {
	c.X1 = max(c.X1, -50)
	c.X2 = min(c.X2, 50)
	c.Y1 = max(c.Y1, -50)
	c.Y2 = min(c.Y2, 50)
	c.Z1 = max(c.Z1, -50)
	c.Z2 = min(c.Z2, 50)
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
