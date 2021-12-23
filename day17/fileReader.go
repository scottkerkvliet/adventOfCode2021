package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadTarget(f string) ([2][2]int, error) {
	var targets [2][2]int

	targetFile, err := ioutil.ReadFile(f)
	if err != nil {
		return targets, fmt.Errorf("Could not open file: %v", f)
	}

	targetString := string(targetFile)[15:]
	targetRanges := strings.Split(targetString, ", y=")
	if len(targetRanges) != 2 {
		return targets, errors.New("Target file does not contain coordinates in expected format")
	}

	xRange := strings.Split(targetRanges[0], "..")
	yRange := strings.Split(targetRanges[1], "..")
	if len(xRange) != 2 || len(yRange) != 2 {
		return targets, errors.New("Target file does not contain coordinates in expected format")
	}

	var err1, err2, err3, err4 error
	targets[0][0], err1 = strconv.Atoi(xRange[0])
	targets[0][1], err2 = strconv.Atoi(xRange[1])
	targets[1][0], err3 = strconv.Atoi(yRange[0])
	targets[1][1], err4 = strconv.Atoi(yRange[1])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return targets, errors.New("One of the coordinates was not a number")
	}

	return targets, nil
}
