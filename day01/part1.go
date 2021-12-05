package main

import (
	"fmt"
)

func main() {
	depths, err := ReadDepths("depths.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numIncreases := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			numIncreases++
		}
	}

	fmt.Printf("There are %v depth increases.\n", numIncreases)
}
