package fileReader

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"errors"
	"Day22/Cuboid"
)

func ReadInstructions(f string) ([]*cuboid.Instruction, error) {
	instructionFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(instructionFile)
	var instructions []*cuboid.Instruction
	for scanner.Scan() {
		instruction, err := getInstruction(scanner.Text())
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func getInstruction(line string) (*cuboid.Instruction, error) {
	on := line[:2] == "on"
	var coordinateLine string
	if on {
		coordinateLine = line[3:]
	} else {
		coordinateLine = line[4:]
	}

	coordinateStrings := strings.Split(coordinateLine, ",")
	if len(coordinateStrings) != 3 {
		return nil, fmt.Errorf("Instruction line does not contain 3 pairs of coordinates: \"%v\"", line)
	}
	x1, x2, err1 := getCoordinates(coordinateStrings[0])
	y1, y2, err2 := getCoordinates(coordinateStrings[1])
	z1, z2, err3 := getCoordinates(coordinateStrings[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, fmt.Errorf("Instruction line has malformed coordinates: \"%v\"", line)
	}

	return cuboid.NewInstruction(x1, x2, y1, y2, z1, z2, on), nil
}

func getCoordinates(line string) (int, int, error) {
	coordinates := strings.Split(line[2:], "..")
	if len(coordinates) != 2 {
		return 0, 0, errors.New("coordinate count")
	}
	c1, err1 := strconv.Atoi(coordinates[0])
	c2, err2 := strconv.Atoi(coordinates[1])
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("coordinate conversion")
	}

	return c1, c2, nil
}
