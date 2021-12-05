package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadDepths(f string) ([]int, error) {
	depthsFile, err := os.Open(f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not open file: %v", f))
	}
	defer depthsFile.Close()

	scanner := bufio.NewScanner(depthsFile)
	var depths []int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, errors.New(fmt.Sprintf("File contained non-integer line: %v", scanner.Text()))
		}
		depths = append(depths, depth)
	}

	return depths, nil
}
