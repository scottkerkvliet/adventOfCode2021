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

func main() {
	cave, err := Cave.ReadCave("cave.txt")
	if err != nil {
		log.Fatal(err)
	}

	printLowestRiskPath(cave)
}
