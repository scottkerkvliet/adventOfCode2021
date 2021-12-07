package main

import (
	"fmt"
	"log"
)

func CountOverlaps(ventMap *[1000][1000]int) {
	overlaps := 0

	for i := range ventMap {
		for j := range ventMap[i] {
			if ventMap[i][j] > 1 {
				overlaps++
			}
		}
	}

	fmt.Printf("There are %v coordinates with 2 or more vents\n", overlaps)
}

func main() {
	ventMap, err := ReadVents("vents.txt", false)
	if err != nil {
		log.Fatal(err)
	}

	CountOverlaps(ventMap)
}
