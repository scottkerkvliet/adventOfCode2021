package main

import (
	"fmt"
	"log"
	"math"
)

func min(a, b int) int {
	min := math.Min(float64(a), float64(b))
	return int(min)
}

func max(a, b int) int {
	max := math.Max(float64(a), float64(b))
	return int(max)
}

func getDiffInQuantity(polymer string, insertions map[byte]map[byte]byte) {
	for i := 0; i < 10; i++ {
		newPolymer := polymer[0:1]
		for j := 1; j < len(polymer); j++ {
			newPolymer += string(insertions[polymer[j-1]][polymer[j]]) + polymer[j:j+1]
		}
		polymer = newPolymer
	}

	charCount := map[rune]int{}
	for _, char := range polymer {
		if _, exists := charCount[char]; !exists {
			charCount[char] = 0
		}
		charCount[char] += 1
	}

	maxCount, minCount := 0, len(polymer)
	for _, count := range charCount {
		maxCount = max(maxCount, count)
		minCount = min(minCount, count)
	}

	fmt.Printf("Difference between most common element and least common element is %v occurrences.\n", maxCount-minCount)
}

func main() {
	polymer, insertions, err := ReadPolymer("polymer.txt")
	if err != nil {
		log.Fatal(err)
	}

	getDiffInQuantity(polymer, insertions)
}
