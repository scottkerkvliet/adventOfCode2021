package main

import (
	"Day12/fileReader"
	"fmt"
	"log"
)

func newPathsToEnd(cave *fileReader.Cave, visited map[string]int, returnedToSmallCave bool) (numPaths int) {
	if cave == nil || visited == nil {
		return
	}
	if cave.Name == "end" {
		return 1
	}

	if !cave.Big {
		if visits, hasBeenVisited := visited[cave.Name]; hasBeenVisited && visits > 0 {
			if cave.Name == "start" || returnedToSmallCave {
				return
			}
			returnedToSmallCave = true
		} else {
			visited[cave.Name] = 0
		}
		visited[cave.Name] += 1
	}

	for _, neighbour := range cave.Connections {
		numPaths += newPathsToEnd(neighbour, visited, returnedToSmallCave)
	}

	if !cave.Big {
		visited[cave.Name] -= 1
	}

	return
}

func printNewNumPaths(start *fileReader.Cave) {
	visited := map[string]int{}
	numPaths := newPathsToEnd(start, visited, false)

	fmt.Printf("There are %v paths to end.\n", numPaths)
}

func main() {
	startCave, err := fileReader.ReadCaves("caves.txt")
	if err != nil {
		log.Fatal(err)
	}

	printNewNumPaths(startCave)
}
