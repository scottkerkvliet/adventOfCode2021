package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadPolymer(f string) (string, map[byte]map[byte]byte, error) {
	polymerFile, err := os.Open(f)
	if err != nil {
		return "", nil, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(polymerFile)
	if !scanner.Scan() {
		return "", nil, errors.New("File was empty")
	}

	polymer := scanner.Text()
	scanner.Scan()
	if len(scanner.Text()) != 0 {
		return "", nil, errors.New("Second line was non-empty")
	}

	insertions := map[byte]map[byte]byte{}
	for scanner.Scan() {
		insertionLine := scanner.Text()
		insertionSteps := strings.Split(insertionLine, " -> ")
		if len(insertionSteps) != 2 || len(insertionSteps[0]) != 2 || len(insertionSteps[1]) != 1 {
			return "", nil, fmt.Errorf("Malformed insertion line: %v", insertionLine)
		}
		if _, exists := insertions[insertionSteps[0][0]]; !exists {
			insertions[insertionSteps[0][0]] = map[byte]byte{}
		}
		insertions[insertionSteps[0][0]][insertionSteps[0][1]] = insertionSteps[1][0]
	}

	return polymer, insertions, nil
}
