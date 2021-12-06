package main

import (
	"fmt"
	"log"
)

func binaryToInt(binary string) (r int) {
	for i := 0; i < len(binary); i++ {
		r = r * 2
		if binary[i] == '1' {
			r += 1
		}
	}

	return
}

func separateDigits(diagnostics []string, i int) (ones []string, zeroes []string) {
	for _, d := range diagnostics {
		if d[i] == '1' {
			ones = append(ones, d)
		} else {
			zeroes = append(zeroes, d)
		}
	}

	return
}

func getOxygen(diagnostics []string) int {
	for i := 0; i < len(diagnostics[0]); i++ {
		ones, zeroes := separateDigits(diagnostics, i)
		if len(ones) >= len(zeroes) {
			diagnostics = ones
		} else {
			diagnostics = zeroes
		}
		if len(diagnostics) == 1 {
			break
		}
	}

	return binaryToInt(diagnostics[0])
}

func getCarbon(diagnostics []string) int {
	for i := 0; i < len(diagnostics[0]); i++ {
		ones, zeroes := separateDigits(diagnostics, i)
		if len(zeroes) <= len(ones) {
			diagnostics = zeroes
		} else {
			diagnostics = ones
		}
		if len(diagnostics) == 1 {
			break
		}
	}

	return binaryToInt(diagnostics[0])
}

func printLifeSupport(diagnostics []string) {
	oxygen := getOxygen(diagnostics)
	carbon := getCarbon(diagnostics)

	fmt.Printf("Oxygen generator rating is %v, and CO2 scrubber rating is %v.\n", oxygen, carbon)
	fmt.Printf("Life support rating is %v.\n", oxygen*carbon)
}

func main() {
	diagnostics, err := ReadDiagnostics("diagnostics.txt")
	if err != nil {
		log.Fatal(err)
	}

	printLifeSupport(diagnostics)
}
