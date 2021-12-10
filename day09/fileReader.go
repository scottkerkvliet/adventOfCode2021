package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadMap(f string) ([][]int, error) {
	heightMap := [][]int{}
	mapFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer mapFile.Close()

	rowLength := -1
	scanner := bufio.NewScanner(mapFile)
	for scanner.Scan() {
		rowString := scanner.Text()
		if rowLength == -1 {
			rowLength = len(rowString)
		} else if rowLength != len(rowString) {
			return nil, fmt.Errorf("Row is not the same length: %v", rowString)
		}

		var row []int
		for i, _ := range rowString {
			num, err := strconv.Atoi(rowString[i : i+1])
			if err != nil {
				return nil, fmt.Errorf("Character is not a number: %v", rowString[i:i+1])
			}
			row = append(row, num)
		}

		heightMap = append(heightMap, row)
	}

	return heightMap, nil
}
