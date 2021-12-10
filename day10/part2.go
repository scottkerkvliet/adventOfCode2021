package main

import (
	"fmt"
	"log"
	"sort"
)

func scoreOpenBraces(openBraces []rune, points map[rune]int) (score int) {
	for i := len(openBraces) - 1; i >= 0; i-- {
		score = score * 5
		score += points[openBraces[i]]
	}
	return
}

func printMedianCompletionSum(syntaxLines []string) {
	var scores []int
	completionPoints := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	matchingBrace := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

lineLoop:
	for _, syntax := range syntaxLines {
		var openBraces []rune
		for _, char := range syntax {
			switch char {
			case '(', '[', '{', '<':
				openBraces = append(openBraces, char)
			case ')', ']', '}', '>':
				if openBraces[len(openBraces)-1] != matchingBrace[char] {
					continue lineLoop
				}
				openBraces = openBraces[:len(openBraces)-1]
			}
		}
		scores = append(scores, scoreOpenBraces(openBraces, completionPoints))
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	median := scores[len(scores)/2]

	fmt.Printf("Completion scores: %v\n", scores)
	fmt.Printf("Median of completion scores is %v.\n", median)
}

func main() {
	syntaxLines, err := ReadSyntax("syntax.txt")
	if err != nil {
		log.Fatal(err)
	}

	printMedianCompletionSum(syntaxLines)
}
