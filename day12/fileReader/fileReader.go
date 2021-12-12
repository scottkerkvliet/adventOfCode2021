package fileReader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	Name        string
	Big         bool
	Connections []*Cave
}

func buildCaveMap(scanner *bufio.Scanner) (*Cave, error) {
	caves := map[string]*Cave{}

	for scanner.Scan() {
		connectionLine := scanner.Text()
		connectedCaves := strings.Split(connectionLine, "-")
		if len(connectedCaves) != 2 {
			return nil, fmt.Errorf("Line does not contain a proper cave connection: %v", connectionLine)
		}
		var caveA, caveB *Cave
		var exists bool

		if caveA, exists = caves[connectedCaves[0]]; !exists {
			caveA = &Cave{connectedCaves[0], connectedCaves[0][0] < 'a', nil}
			caves[connectedCaves[0]] = caveA
		}
		if caveB, exists = caves[connectedCaves[1]]; !exists {
			caveB = &Cave{connectedCaves[1], connectedCaves[1][0] < 'a', nil}
			caves[connectedCaves[1]] = caveB
		}

		caveA.Connections = append(caveA.Connections, caveB)
		caveB.Connections = append(caveB.Connections, caveA)
	}

	if _, exists := caves["start"]; !exists {
		return nil, errors.New("No \"start\" cave")
	}
	if _, exists := caves["end"]; !exists {
		return nil, errors.New("No \"end\" cave")
	}

	return caves["start"], nil
}

func ReadCaves(f string) (*Cave, error) {
	caveFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer caveFile.Close()

	return buildCaveMap(bufio.NewScanner(caveFile))
}
