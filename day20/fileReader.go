package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadImage(f string) (string, [][]byte, error) {
	imageFile, err := os.Open(f)
	if err != nil {
		return "", nil, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(imageFile)
	if !scanner.Scan() {
		return "", nil, fmt.Errorf("File was empty: %v", f)
	}
	algorithm := scanner.Text()
	if len(algorithm) != 512 {
		return "", nil, errors.New("Algorithm string was not 512 characters")
	}

	var image [][]byte
	scanner.Scan()
	for scanner.Scan() {
		image = append(image, []byte(scanner.Text()))
	}

	return algorithm, image, nil
}
