package main

import (
	"Day12/fileReader"
	"fmt"
	"log"
)

func main() {
	startCave, err := fileReader.ReadCaves("caves.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Start cave has key \"%v\". Is big? %v\n", startCave.Name, startCave.Big)
	fmt.Printf("Start cave has %v connections.\n", len(startCave.Connections))
}
