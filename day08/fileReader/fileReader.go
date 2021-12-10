package fileReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Digit map[rune]bool

type Display struct {
	Digits  []Digit
	Outputs []Digit
}

func makeDigit(runes string) Digit {
	digit := make(Digit)
	for _, rune := range runes {
		digit[rune] = true
	}

	return digit
}

func getDisplay(displayString string) (Display, error) {
	display := Display{}
	dislayStringSplit := strings.Split(displayString, " | ")
	if len(dislayStringSplit) != 2 {
		return display, fmt.Errorf("Display line does not contain proper digits/output format: %v", displayString)
	}

	digits := strings.Split(dislayStringSplit[0], " ")
	for _, digitString := range digits {
		display.Digits = append(display.Digits, makeDigit(digitString))
	}
	outputs := strings.Split(dislayStringSplit[1], " ")
	for _, outputString := range outputs {
		display.Outputs = append(display.Outputs, makeDigit(outputString))
	}

	return display, nil
}

func ReadDisplays(f string) ([]Display, error) {
	displaysFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer displaysFile.Close()

	var displays []Display
	scanner := bufio.NewScanner(displaysFile)
	for scanner.Scan() {
		display, err := getDisplay(scanner.Text())
		if err != nil {
			return nil, err
		}
		displays = append(displays, display)
	}

	return displays, nil
}
