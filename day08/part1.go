package main

import (
	"Day8/fileReader"
	"fmt"
	"log"
)

func countOutputDigits(displays []fileReader.Display) {
	count := 0
	for _, display := range displays {
		for _, output := range display.Outputs {
			segments := len(output)
			if segments == 2 || segments == 4 || segments == 3 || segments == 7 {
				count++
			}
		}
	}

	fmt.Printf("There are %v output digits that are 1, 4, 7, or 8.\n", count)
}

func main() {
	displays, err := fileReader.ReadDisplays("displays.txt")
	if err != nil {
		log.Fatal(err)
	}

	countOutputDigits(displays)
}
