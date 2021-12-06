package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDiagnostics(f string) ([]string, error) {
	diagnosticsFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer diagnosticsFile.Close()

	scanner := bufio.NewScanner(diagnosticsFile)
	var diagnostics []string
	for scanner.Scan() {
		diagnostic := scanner.Text()
		for _, c := range diagnostic {
			if c != '0' && c != '1' {
				return nil, fmt.Errorf("File contained non-integer line: %v", diagnostic)
			}
		}
		diagnostics = append(diagnostics, diagnostic)
	}

	return diagnostics, nil
}
