package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSyntax(f string) ([]string, error) {
	syntaxFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file %v", f)
	}
	defer syntaxFile.Close()

	var syntaxLines []string
	scanner := bufio.NewScanner(syntaxFile)
	for scanner.Scan() {
		syntaxLines = append(syntaxLines, scanner.Text())
	}

	return syntaxLines, nil
}
