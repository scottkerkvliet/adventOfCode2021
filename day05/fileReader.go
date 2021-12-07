package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func getPoint(pointString string) (*Point, error) {
	coordStrings := strings.Split(pointString, ",")
	if len(coordStrings) != 2 {
		return nil, fmt.Errorf("Point does not contain 2 coordinates: \"%v\"", pointString)
	}

	coordX, err := strconv.Atoi(coordStrings[0])
	if err != nil {
		return nil, fmt.Errorf("Coordinate is not an integer: %v", coordStrings[0])
	}
	coordY, err := strconv.Atoi(coordStrings[1])
	if err != nil {
		return nil, fmt.Errorf("Coordinate is not an integer: %v", coordStrings[1])
	}

	return &Point{coordX, coordY}, nil
}

func getPoints(ventString string) (*Point, *Point, error) {
	pointStrings := strings.Split(ventString, " -> ")
	if len(pointStrings) != 2 {
		return nil, nil, fmt.Errorf("Line does not contain 2 points: \"%v\"", ventString)
	}

	var point1, point2 *Point
	point1, err := getPoint(pointStrings[0])
	if err == nil {
		point2, err = getPoint(pointStrings[1])
	}
	if err != nil {
		return nil, nil, err
	}

	return point1, point2, nil
}

func addVentToMap(ventString string, ventMap *[1000][1000]int, diagonals bool) error {
	point1, point2, err := getPoints(ventString)
	if err != nil {
		return err
	}

	if !diagonals && point1.X != point2.X && point1.Y != point2.Y {
		return nil
	}

	var xDirection, yDirection, distance int
	switch {
	case point1.X < point2.X:
		xDirection = 1
		distance = point2.X - point1.X + 1
	case point1.X > point2.X:
		xDirection = -1
		distance = point1.X - point2.X + 1
	default:
		xDirection = 0
	}
	switch {
	case point1.Y < point2.Y:
		yDirection = 1
		distance = point2.Y - point1.Y + 1
	case point1.Y > point2.Y:
		yDirection = -1
		distance = point1.Y - point2.Y + 1
	default:
		yDirection = 0
	}

	for i := 0; i < distance; i++ {
		ventMap[point1.X+(xDirection*i)][point1.Y+(yDirection*i)] += 1
	}

	return nil
}

func ReadVents(f string, diagonals bool) (*[1000][1000]int, error) {
	ventMap := [1000][1000]int{}
	ventsFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer ventsFile.Close()

	scanner := bufio.NewScanner(ventsFile)
	for scanner.Scan() {
		err := addVentToMap(scanner.Text(), &ventMap, diagonals)
		if err != nil {
			return nil, err
		}
	}

	return &ventMap, nil
}
