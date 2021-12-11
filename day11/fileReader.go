package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func ReadOctopuses(f string) ([][]int, error) {
	octopusFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file %v", f)
	}

	var octopuses [][]int
	rowLength := -1
	scanner := bufio.NewScanner(octopusFile)
	for scanner.Scan() {
	  row := scanner.Text()
		if rowLength == -1 {
			rowLength = len(row)
		} else if len(row) != rowLength {
			return nil, fmt.Errorf("Row is not the same length: %v", row)
		}

		var octopusRow []int
		for i, _ := range row {
			octopus, err := strconv.Atoi(row[i:i+1])
			if err != nil {
				fmt.Errorf("Character is not a number: %v", row[i:i+1])
			}
			octopusRow = append(octopusRow, octopus)
		}
		octopuses = append(octopuses, octopusRow)
	}

	return octopuses, nil
}
