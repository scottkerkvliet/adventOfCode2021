package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCrabSlice(allCrabs string) ([]int, error) {
	crabStrings := strings.Split(allCrabs, ",")
	crabs := make([]int, len(crabStrings))

	for i, crabString := range crabStrings {
		crab, err := strconv.Atoi(crabString)
		if err != nil {
			return nil, fmt.Errorf("Item not a number: \"%v\"", crabString)
		}
		crabs[i] = crab
	}

	return crabs, nil
}

func ReadCrabs(f string) ([]int, error) {
	crabFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer crabFile.Close()

	scanner := bufio.NewScanner(crabFile)
	if !scanner.Scan() {
		return nil, fmt.Errorf("File is empty: %v", f)
	}
	crabSlice, err := getCrabSlice(scanner.Text())
	if err != nil {
		return nil, err
	}

	return crabSlice, nil
}
