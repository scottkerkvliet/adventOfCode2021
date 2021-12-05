package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Down
	Up
)

type Command struct {
	Direction Direction
	Length    int
}

func ReadCommands(f string) ([]Command, error) {
	commandsFile, err := os.Open(f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not open file: %v", f))
	}
	defer commandsFile.Close()

	scanner := bufio.NewScanner(commandsFile)
	var commands []Command

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var direction Direction
		var length int

		validLine := false
		if len(line) == 2 {
			length, err = strconv.Atoi(line[1])
			if err == nil {
				validLine = true
				switch line[0] {
				case "forward":
					direction = Forward
				case "down":
					direction = Down
				case "up":
					direction = Up
				default:
					validLine = false
				}
			}
		}

		if !validLine {
			return nil, errors.New(fmt.Sprintf("File contained invalid line: %v", scanner.Text()))
		}
		commands = append(commands, Command{direction, length})
	}

	return commands, nil
}
