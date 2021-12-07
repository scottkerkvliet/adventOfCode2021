package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetBingoNumbers(fileNums string) ([]int, error) {
	var nums []int
	numStrings := strings.Split(fileNums, ",")

	for _, numString := range numStrings {
		newNum, err := strconv.Atoi(numString)
		if err != nil {
			return nil, fmt.Errorf("Bingo number not a number: \"%v\"", numString)
		}
		nums = append(nums, newNum)
	}

	return nums, nil
}

func GetBingoBoard(rawBoard [5]string) ([5][5]int, error) {
	board := [5][5]int{}

	for i := 0; i < 5; i++ {
		if len(rawBoard[i]) < 13 {
			return board, fmt.Errorf("Invalid board row: \"%v\"", rawBoard[i])
		}

		for j := 0; j < 5; j++ {
			numString := rawBoard[i][j*3 : (j*3)+2]
			if numString[0] == ' ' {
				numString = numString[1:2]
			}
			num, err := strconv.Atoi(numString)
			if err != nil {
				return board, fmt.Errorf("Invalid board number: \"%v\"", numString)
			}
			board[i][j] = num
		}
	}

	return board, nil
}

func ReadBingo(f string) ([]int, [][5][5]int, error) {
	bingoFile, err := os.Open(f)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open file: %v", f)
	}
	defer bingoFile.Close()

	scanner := bufio.NewScanner(bingoFile)

	// Get numbers
	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("File was empty: %v", f)
	}
	nums, err := GetBingoNumbers(scanner.Text())
	if err != nil {
		return nil, nil, err
	}

	// Get boards
	var boards [][5][5]int
	for scanner.Scan() {
		rawBoard := [5]string{}
		for i := 0; i < 5; i++ {
			if !scanner.Scan() {
				return nil, nil, fmt.Errorf("Malformed board")
			}
			rawBoard[i] = scanner.Text()
		}

		board, err := GetBingoBoard(rawBoard)
		if err != nil {
			return nil, nil, err
		}
		boards = append(boards, board)
	}

	return nums, boards, nil
}
