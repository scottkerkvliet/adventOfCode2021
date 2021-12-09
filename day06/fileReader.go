package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFishSlice(allFish string) (*[]int, error) {
	fishStrings := strings.Split(allFish, ",")
	fish := make([]int, len(fishStrings))

	for i, fishString := range fishStrings {
		fishNum, err := strconv.Atoi(fishString)
		if err != nil {
			return nil, fmt.Errorf("Item not a number: \"%v\"", fishString)
		}
		fish[i] = fishNum
	}

	return &fish, nil
}

func ReadFish(f string) (*[]int, error) {
	fishFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer fishFile.Close()

	scanner := bufio.NewScanner(fishFile)
	if !scanner.Scan() {
		return nil, fmt.Errorf("File is empty: %v", f)
	}
	fishSlice, err := getFishSlice(scanner.Text())
	if err != nil {
		return nil, err
	}

	return fishSlice, nil
}
