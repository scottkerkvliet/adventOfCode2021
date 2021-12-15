package Cave

import (
	"fmt"
	"os"
	"bufio"
)

func ReadCave(f string) ([][]int, error) {
	caveFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(caveFile)
	var cave [][]int
	for scanner.Scan() {
		rowString := scanner.Text()
		var caveRow []int
		for _, char := range rowString {
			caveRow = append(caveRow, int(char - '0'))
		}
		cave = append(cave, caveRow)
	}

	return cave, nil
}
