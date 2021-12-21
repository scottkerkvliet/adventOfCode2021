package filereader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readPositionLine(line string) (int, error) {
	if len(line) < 29 {
		return 0, fmt.Errorf("Line not in expected format: \"%v\"", line)
	}
	position, err := strconv.Atoi(line[28:])
	if err != nil {
		return 0, fmt.Errorf("Position is not a number: \"%v\"", line[28:])
	}

	return position, nil
}

func ReadPositions(f string) (int, int, error) {
	positionFile, err := os.Open(f)
	if err != nil {
		return 0, 0, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(positionFile)
	if !scanner.Scan() {
		return 0, 0, errors.New("File did not contain enough lines")
	}
	player1Position, err := readPositionLine(scanner.Text())
	if err != nil {
		return 0, 0, err
	}
	if !scanner.Scan() {
		return 0, 0, errors.New("File did not contain enough lines")
	}
	player2Position, err := readPositionLine(scanner.Text())
	if err != nil {
		return 0, 0, err
	}

	return player1Position, player2Position, nil
}
