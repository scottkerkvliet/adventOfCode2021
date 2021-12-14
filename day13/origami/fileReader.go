package origami

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func getPoint(pointLine string) (Point, error) {
	pointStrings := strings.Split(pointLine, ",")
	if len(pointStrings) != 2 {
		return Point{}, fmt.Errorf("Invalid coordinates: \"%v\"", pointLine)
	}

	x, err := strconv.Atoi(pointStrings[0])
	if err != nil {
		return Point{}, fmt.Errorf("Invalid x coordinate: \"%v\"", pointStrings[0])
	}
	y, err := strconv.Atoi(pointStrings[1])
	if err != nil {
		return Point{}, fmt.Errorf("Invalid y coordinate: \"%v\"", pointStrings[1])
	}

	return Point{x, y}, nil
}

func getFold(foldLine string) (Fold, error) {
	if len(foldLine) < 14 {
		return Fold{}, fmt.Errorf("Invalid fold line: \"%v\"", foldLine)
	}
	xDirection := foldLine[11] == 'x'
	foldLocation, err := strconv.Atoi(foldLine[13:])
	if err != nil {
		return Fold{}, fmt.Errorf("Invalid fold location: \"%v\"", foldLine[13:])
	}

	return Fold{foldLocation, xDirection}, nil
}

func ReadOrigami(f string) (*Origami, []Fold, error) {
	origamiFile, err := os.Open(f)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open file: %v", f)
	}

	scanner := bufio.NewScanner(origamiFile)
	var points []Point
	for scanner.Scan() {
		pointLine := scanner.Text()
		if len(pointLine) == 0 {
			break
		}

		point, err := getPoint(pointLine)
		if err != nil {
			return nil, nil, err
		}

		points = append(points, point)
	}

	origami := &Origami{points}

	var folds []Fold
	for scanner.Scan() {
		foldLine := scanner.Text()
		fold, err := getFold(foldLine)
		if err != nil {
			return nil, nil, err
		}

		folds = append(folds, fold)
	}

	return origami, folds, nil
}
