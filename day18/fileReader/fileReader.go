package filereader

import (
	sn "Day18/snailNumber"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readNumberLine(line string) (sn.SnailNumber, error) {
	var numbers []sn.SnailNumber

	for i, char := range line {
		switch char {
		case '[', ',':
			break
		case ']':
			if len(numbers) < 2 {
				return nil, fmt.Errorf("Closing brace came before too early in line: \"%v\"", line[:i+1])
			}
			newNumber := sn.NewPair(numbers[len(numbers)-2], numbers[len(numbers)-1])
			numbers = append(numbers[:len(numbers)-2], newNumber)
		default:
			value, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("Invalid character: \"%v\"", char)
			}
			numbers = append(numbers, sn.NewSingle(value))
		}
	}

	if len(numbers) != 1 {
		return nil, fmt.Errorf("Invalid line, not exactly 1 number: \"%v\"", line)
	}

	return numbers[0], nil
}

func ReadNumbers(f string) ([]sn.SnailNumber, error) {
	numbersFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer numbersFile.Close()

	var numbers []sn.SnailNumber
	scanner := bufio.NewScanner(numbersFile)
	for scanner.Scan() {
		number, err := readNumberLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}
