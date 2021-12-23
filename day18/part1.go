package main

import (
	fr "Day18/fileReader"
	sn "Day18/snailNumber"
	"fmt"
	"log"
)

func main() {
	numbers, err := fr.ReadNumbers("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalNumber := numbers[0]
	for i := 1; i < len(numbers); i++ {
		finalNumber = sn.Add(finalNumber, numbers[i])
	}

	fmt.Printf("The final number is: %v\n", finalNumber.Print())
	fmt.Printf("The magnitude is %v\n", finalNumber.GetMagnitude())
}
