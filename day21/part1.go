package main

import (
	fr "Day21/FileReader"
	game "Day21/Game"
	"fmt"
	"log"
)

func main() {
	const maxScore = 1000
	player1Position, player2Position, err := fr.ReadPositions("positions.txt")
	if err != nil {
		log.Fatal(err)
	}

	die := game.NewDeterministicDie()
	player1 := game.NewPlayer(player1Position)
	player2 := game.NewPlayer(player2Position)

	for player1.Score < maxScore && player2.Score < maxScore {
		player1 = game.MovePlayer(player1, die)
		if player1.Score < maxScore {
			player2 = game.MovePlayer(player2, die)
		}
	}

	lowestScore := player2.Score
	if player1.Score < maxScore {
		lowestScore = player1.Score
	}
	fmt.Printf("The lowest score was %v and the number of dice rolls was %v.\n", lowestScore, die.GetCount())
	fmt.Printf("The product of these values is %v.\n", lowestScore*die.GetCount())
}
