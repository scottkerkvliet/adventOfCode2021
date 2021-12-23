package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func min(a, b int) int {
	min := math.Min(float64(a), float64(b))
	return int(min)
}

func max(a, b int) int {
	max := math.Max(float64(a), float64(b))
	return int(max)
}

func GetMaxMinCounts(cc map[byte]int) (int, int) {
	maxCount, minCount := 0, math.MaxInt
	for _, count := range cc {
		maxCount = max(maxCount, count)
		minCount = min(minCount, count)
	}

	return maxCount, minCount
}

func MergeMaps(a, b map[byte]int) {
	for bKey, bCount := range b {
		if _, exists := a[bKey]; !exists {
			a[bKey] = 0
		}
		a[bKey] += bCount
	}
}

func countCharsToDepth(a, b byte, insertions map[byte]map[byte]byte, currentCount map[byte]int, depth int, precomputeCounts map[byte]map[byte]map[byte]int) {
	middleChar := insertions[a][b]
	if _, exists := currentCount[middleChar]; !exists {
		currentCount[middleChar] = 0
	}
	currentCount[middleChar]++
	if depth > 1 {
		countCharsToDepth(a, middleChar, insertions, currentCount, depth-1, precomputeCounts)
		countCharsToDepth(middleChar, b, insertions, currentCount, depth-1, precomputeCounts)
	} else if precomputeCounts != nil {
		MergeMaps(currentCount, precomputeCounts[a][middleChar])
		MergeMaps(currentCount, precomputeCounts[middleChar][b])
	}
}

func precomputePair(a, b byte, insertions map[byte]map[byte]byte, depth int) map[byte]int {
	finalCount := map[byte]int{}
	countCharsToDepth(a, b, insertions, finalCount, depth, nil)
	return finalCount
}

func getDiffInQuantity(polymer string, insertions map[byte]map[byte]byte) {
	totalDepth := 40
	precomputeDepth := 20

	start := time.Now()
	precomputeCounts := map[byte]map[byte]map[byte]int{}
	for char1, char1Map := range insertions {
		if _, exists := precomputeCounts[char1]; !exists {
			precomputeCounts[char1] = map[byte]map[byte]int{}
		}
		for char2, _ := range char1Map {
			precomputeCounts[char1][char2] = precomputePair(char1, char2, insertions, precomputeDepth)
		}
	}
	fmt.Printf("Finished Precompute      (Time Elapsed: %v)\n", time.Since(start))

	charCount := map[byte]int{}
	charCount[polymer[0]] = 1
	for i := 1; i < len(polymer); i++ {
		if _, exists := charCount[polymer[i]]; !exists {
			charCount[polymer[i]] = 0
		}
		charCount[polymer[i]]++

		countCharsToDepth(polymer[i-1], polymer[i], insertions, charCount, totalDepth-precomputeDepth, precomputeCounts)
		fmt.Printf("Char #%v       (Time Elapsed: %v)\n", i, time.Since(start))
	}

	maxCount, minCount := GetMaxMinCounts(charCount)
	fmt.Printf("Difference between most common element and least common element is %v occurrences.\n", maxCount-minCount)
}

func main() {
	polymer, insertions, err := ReadPolymer("polymer.txt")
	if err != nil {
		log.Fatal(err)
	}

	getDiffInQuantity(polymer, insertions)
}
