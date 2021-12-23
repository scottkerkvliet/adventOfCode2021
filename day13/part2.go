package main

import (
	"Day13/origami"
	"fmt"
	"log"
	"math"
)

func max(a, b int) int {
	max := math.Max(float64(a), float64(b))

	return int(max)
}

func printGrid(grid [][]bool) {
	chars := map[bool]string{
		true: "#",
		false: " ",
	}

	for _, row := range grid {
		for _, point := range row {
			fmt.Print(chars[point])
		}
		fmt.Print("\n")
	}
}

func printOrigami(o *origami.Origami) {
	var grid [][]bool

	var maxX, maxY int
	for _, point := range o.Points {
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
	}

	for i := 0; i <= maxY; i++ {
		grid = append(grid, make([]bool, maxX+1))
	}

	for _, point := range o.Points {
		grid[point.Y][point.X] = true
	}

	printGrid(grid)
}

func main() {
	origami, folds, err := origami.ReadOrigami("origami.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, fold := range folds {
		origami.Fold(fold)
	}

	printOrigami(origami)
}
