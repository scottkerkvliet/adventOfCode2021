package main

import (
	"fmt"
	"log"
)

func printSumSyntaxErrors(syntaxLines []string) {
	sum := 0
	errorPoints := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	matchingBrace := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	for _, syntax := range syntaxLines {
		var openBraces []rune
	syntaxLoop:
		for _, char := range syntax {
			switch char {
			case '(', '[', '{', '<':
				openBraces = append(openBraces, char)
			case ')', ']', '}', '>':
				if openBraces[len(openBraces)-1] != matchingBrace[char] {
					sum += errorPoints[char]
					break syntaxLoop
				}
				openBraces = openBraces[:len(openBraces)-1]
			}
		}
	}

	fmt.Printf("Sum of syntax errors is %v.\n", sum)
}

func main() {
	syntaxLines, err := ReadSyntax("syntax.txt")
	if err != nil {
		log.Fatal(err)
	}

	printSumSyntaxErrors(syntaxLines)
}
