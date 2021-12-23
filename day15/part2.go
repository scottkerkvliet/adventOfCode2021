package main

import (
	"Day15/Cave"
	"container/heap"
	"fmt"
	"log"
)

func printLowestRiskPath(cave [][]int) {
	numCols, numRows := len(cave[0]), len(cave)
	var lowestRiskToEnd [][]int
	for i := 0; i < numRows; i++ {
		lowestRiskToEnd = append(lowestRiskToEnd, make([]int, numCols))
	}

	priorityQueue := make(Cave.PriorityQueue, 1)
	priorityQueue[0] = Cave.MakePoint(numCols-1, numRows-1, 0)
	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		point := heap.Pop(&priorityQueue).(*Cave.Point)
		if lowestRiskToEnd[point.Y][point.X] != 0 {
			continue
		}
		lowestRiskToEnd[point.Y][point.X] = point.Risk
		if point.Y == 0 && point.X == 0 {
			break
		}

		if point.Y > 0 && lowestRiskToEnd[point.Y-1][point.X] == 0 {
			heap.Push(&priorityQueue, Cave.MakePoint(point.X, point.Y-1, point.Risk+cave[point.Y][point.X]))
		}
		if point.Y+1 < numRows && lowestRiskToEnd[point.Y+1][point.X] == 0 {
			heap.Push(&priorityQueue, Cave.MakePoint(point.X, point.Y+1, point.Risk+cave[point.Y][point.X]))
		}
		if point.X > 0 && lowestRiskToEnd[point.Y][point.X-1] == 0 {
			heap.Push(&priorityQueue, Cave.MakePoint(point.X-1, point.Y, point.Risk+cave[point.Y][point.X]))
		}
		if point.X+1 < numCols && lowestRiskToEnd[point.Y][point.X+1] == 0 {
			heap.Push(&priorityQueue, Cave.MakePoint(point.X+1, point.Y, point.Risk+cave[point.Y][point.X]))
		}
	}

	fmt.Printf("The lowest risk path has a risk level of %v.\n", lowestRiskToEnd[0][0])
}

func addRiskToCave(cave [][]int, addedRisk int) (newCave [][]int) {
	for i := 0; i < len(cave); i++ {
		newCave = append(newCave, make([]int, len(cave[i])))
		for j := 0; j < len(cave[i]); j++ {
			newCave[i][j] = cave[i][j] + addedRisk
			for newCave[i][j] > 9 {
				newCave[i][j] -= 9
			}
		}
	}

	return newCave
}

func generateRealCave(cave [][]int) (realCave [][]int) {
	multiplyBy := 5
	numRows, numCols := len(cave), len(cave[0])

	for i := 0; i < numRows*multiplyBy; i++ {
		realCave = append(realCave, make([]int, numCols*multiplyBy))
	}

	referenceCaves := map[int][][]int{}

	for i := 0; i < multiplyBy; i++ {
		for j := 0; j < multiplyBy; j++ {
			referenceCave, exists := referenceCaves[i+j]
			if i == 0 && j == 0 {
				referenceCave = cave
			} else if !exists {
				referenceCave = addRiskToCave(cave, i+j)
				referenceCaves[i+j] = referenceCave
			}

			for r, row := range referenceCave {
				copy(realCave[i*numRows+r][j*numCols:], row)
			}
		}
	}

	return realCave
}

func main() {
	cave, err := Cave.ReadCave("cave.txt")
	if err != nil {
		log.Fatal(err)
	}

	realCave := generateRealCave(cave)
	printLowestRiskPath(realCave)
}
