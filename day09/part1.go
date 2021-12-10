package main

import (
	"fmt"
	"log"
)

func printRiskLevel(heightMap [][]int) {
	riskLevel := 0
	rowLength := len(heightMap)
	colLength := len(heightMap[0])

	for row, _ := range heightMap {
		for col, _ := range heightMap {
			switch {
			case row > 0 && heightMap[row][col] >= heightMap[row-1][col]:
				continue
			case row+1 < rowLength && heightMap[row][col] >= heightMap[row+1][col]:
				continue
			case col > 0 && heightMap[row][col] >= heightMap[row][col-1]:
				continue
			case col+1 < colLength && heightMap[row][col] >= heightMap[row][col+1]:
				continue
			default:
				riskLevel += heightMap[row][col] + 1
			}
		}
	}

	fmt.Printf("Risk level is %v.\n", riskLevel)
}

func main() {
	heightMap, err := ReadMap("map.txt")
	if err != nil {
		log.Fatal(err)
	}

	printRiskLevel(heightMap)
}
