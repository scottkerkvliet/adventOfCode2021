package main

import (
	"Day24/alu"
	fr "Day24/fileReader"
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	instructions, err := fr.ReadInstructions("monad.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	ALU := alu.NewALU()
	inputLength := 14
	input := 0
	for i := 0; i < inputLength; i++ {
		input = (input * 10) + 9
	}

	for ; input > 0; input-- {
		elapsed := time.Since(start)
		if elapsed.Seconds() > 5 {
			fmt.Printf("input %v\n", input)
			start = time.Now()
		}
		inputs := make(chan int, inputLength)
		if !populateChannelWithX(inputs, input, inputLength) {
			continue
		}
		close(inputs)
		err = ALU.RunProgram(instructions, inputs)
		if err != nil {
			log.Fatal(err)
		}

		if ALU.GetZ() == 1 {
			fmt.Printf("Found valid number: %v\n", input)
			break
		}
	}
}

func populateChannelWithX(inputs chan int, input int, num int) bool {
	for num > 0 {
		num--
		factor := int(math.Pow(10, float64(num)))
		digit := num / factor
		if digit == 0 {
			return false
		}
		inputs <- digit
		num -= digit * factor
	}

	return true
}
