package main

import (
	"Day12/fileReader"
	"fmt"
	"log"
)

func pathsToEnd(cave *fileReader.Cave, visited map[string]bool) (numPaths int) {
	if cave == nil || visited == nil {
		return
	}
	if cave.Name == "end" {
		return 1
	}
	if !cave.Big {
		if _, hasBeenVisited := visited[cave.Name]; hasBeenVisited {
			return
		}
		visited[cave.Name] = true
	}

	for _, neighbour := range cave.Connections {
		numPaths += pathsToEnd(neighbour, visited)
	}

	if !cave.Big {
		delete(visited, cave.Name)
	}

	return
}

func printNumPaths(start *fileReader.Cave) {
	visitedCaves := map[string]bool{}
	numPaths := pathsToEnd(start, visitedCaves)

	fmt.Printf("There are %v paths to end.\n", numPaths)
}

func main() {
	startCave, err := fileReader.ReadCaves("caves.txt")
	if err != nil {
		log.Fatal(err)
	}

	printNumPaths(startCave)
}
