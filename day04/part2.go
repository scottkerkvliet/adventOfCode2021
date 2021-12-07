package main

import (
	"fmt"
	"log"
)

func checkBoardAt(board *[5][5]int, row, col int) bool {
	rowLine := true
	colLine := true
	for i := 0; i < 5; i++ {
		rowLine = rowLine && board[row][i] == -1
		colLine = colLine && board[i][col] == -1
	}

	return rowLine || colLine
}

func markNumberOnBoard(num int, board *[5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == num {
				board[i][j] = -1
				return checkBoardAt(board, i, j)
			}
		}
	}

	return false
}

func getBoardScore(board [5][5]int) (sum int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}

	return
}

func printLastBingoScore(nums []int, boards [][5][5]int) {
	for _, num := range nums {
		k := 0
		for i := range boards {
			hasLine := markNumberOnBoard(num, &boards[i])
			if hasLine {
				if len(boards) == 1 {
					boardScore := getBoardScore(boards[i])
					fmt.Printf("Last board won when %v was called. Board score is %v.\n", num, boardScore)
					fmt.Printf("Total score is %v.\n", num*boardScore)
					return
				}

			} else {
				if i != k {
					boards[k] = boards[i]
				}
				k++
			}
		}
		boards = boards[:k]
	}

	fmt.Println("All numbers used.")
}

func main() {
	nums, boards, err := ReadBingo("bingo.txt")
	if err != nil {
		log.Fatal(err)
	}

	printLastBingoScore(nums, boards)
}
