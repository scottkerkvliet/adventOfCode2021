package main

import (
	"fmt"
	"log"
)

var keepLargest = 3

func visitBasin(row, col int, heightMap [][]int, basinsVisited [][]bool) (size int) {
	if row < 0 || col < 0 || row >= len(basinsVisited) || col >= len(basinsVisited) || basinsVisited[row][col] || heightMap[row][col] == 9 {
		return
	}
	size++
	basinsVisited[row][col] = true
	size += visitBasin(row-1, col, heightMap, basinsVisited)
	size += visitBasin(row+1, col, heightMap, basinsVisited)
	size += visitBasin(row, col-1, heightMap, basinsVisited)
	size += visitBasin(row, col+1, heightMap, basinsVisited)

	return size
}

func updateLargest(largest []int, num int) {
	for i := 0; i < len(largest); i++ {
		if num < largest[i] {
			return
		}
		if i > 0 {
			largest[i-1] = largest[i]
		}
		largest[i] = num
	}
}

func printProductBasins(heightMap [][]int) {
	rowLength := len(heightMap)
	colLength := len(heightMap[0])
	largestBasins := make([]int, keepLargest)
	var basinsVisited [][]bool

	for i := 0; i < rowLength; i++ {
		basinsVisited = append(basinsVisited, make([]bool, colLength))
	}

	for row, _ := range heightMap {
		for col, _ := range heightMap {
			switch {
			case basinsVisited[row][col]:
				continue
			case row > 0 && heightMap[row][col] >= heightMap[row-1][col]:
				continue
			case row+1 < rowLength && heightMap[row][col] >= heightMap[row+1][col]:
				continue
			case col > 0 && heightMap[row][col] >= heightMap[row][col-1]:
				continue
			case col+1 < colLength && heightMap[row][col] >= heightMap[row][col+1]:
				continue
			default:
				size := visitBasin(row, col, heightMap, basinsVisited)
				updateLargest(largestBasins, size)
			}
		}
	}

	product := 1
	for _, basin := range largestBasins {
		product = product * basin
	}

	fmt.Printf("The product of the %v largest basins is %v.\n", keepLargest, product)
}

func main() {
	heightMap, err := ReadMap("map.txt")
	if err != nil {
		log.Fatal(err)
	}

	printProductBasins(heightMap)
}
