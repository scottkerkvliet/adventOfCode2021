package main

import (
	"Day8/fileReader"
	"fmt"
	"log"
	"reflect"
)

var cycles = 0

func main() {
	displays, err := fileReader.ReadDisplays("displays.txt")
	if err != nil {
		log.Fatal(err)
	}

	printOutputSum(displays)
}

func printOutputSum(displays []fileReader.Display) {
	sum := 0
	for _, display := range displays {
		sum += getOutput(display)
	}

	fmt.Printf("The sum of the outputs is %v.\n", sum)
}

func getOutput(display fileReader.Display) int {
	digitReference := determineDigits(display)
	output := 0
	for _, outputDigit := range display.Outputs {
		for i, digit := range digitReference {
			if reflect.DeepEqual(digit, outputDigit) {
				output = (output * 10) + i
				break
			}
		}
	}

	return output
}

//***************************************************** Determining Digits *******************************************************************

func determineDigits(display fileReader.Display) (digits [10]fileReader.Digit) {
	digits[1] = findDigitWithAll(display.Digits, 2, nil)
	digits[4] = findDigitWithAll(display.Digits, 4, nil)
	digits[7] = findDigitWithAll(display.Digits, 3, nil)
	digits[8] = findDigitWithAll(display.Digits, 7, nil)

	digits[3] = findDigitWithAll(display.Digits, 5, getUniqueRunes(digits[1], nil))
	digits[9] = findDigitWithAll(display.Digits, 6, getUniqueRunes(digits[3], nil))
	digits[6] = findDigitMissingSome(display.Digits, 6, getUniqueRunes(digits[1], nil))
	digits[0] = findDigitMissingSome(display.Digits, 6, getUniqueRunes(digits[4], digits[1]))
	digits[5] = findDigitWithAll(display.Digits, 5, getUniqueRunes(digits[9], digits[3]))
	digits[2] = findDigitWithAll(display.Digits, 5, getUniqueRunes(digits[6], digits[5]))

	return
}

//******************************************************* Finding Digits *******************************************************************

func findDigitMissingSome(digits []fileReader.Digit, length int, runes []rune) fileReader.Digit {
	for _, digit := range digits {
		if len(digit) != length {
			continue
		}
		for _, rune := range runes {
			_, exists := digit[rune]
			if !exists {
				return digit
			}
		}
	}

	return nil
}

func findDigitWithAll(digits []fileReader.Digit, length int, runes []rune) fileReader.Digit {
	for _, digit := range digits {
		if len(digit) != length {
			continue
		}

		match := true
		for _, rune := range runes {
			_, exists := digit[rune]
			if !exists {
				match = false
				break
			}
		}

		if match {
			return digit
		}
	}

	return nil
}

func getUniqueRunes(uniqueDigit fileReader.Digit, baseDigit fileReader.Digit) (runes []rune) {
	for rune, _ := range uniqueDigit {
		_, exists := baseDigit[rune]
		if !exists {
			runes = append(runes, rune)
		}
	}

	return runes
}
