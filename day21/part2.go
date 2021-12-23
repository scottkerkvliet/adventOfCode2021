package main

import (
	fr "Day21/FileReader"
	game "Day21/Game"
	"fmt"
	"log"
)

const maxScore = 21

type Universe struct {
	player1 game.Player
	player2 game.Player
	count   int
}

func processUniverse(player1, player2 game.Player, count int) (player1Wins int, player2Wins int) {
	die1 := game.NewQuantumDie()
	for !die1.IsComplete() {
		roll1Die, roll1Universes := die1.GetDie()
		newPlayer1 := game.MovePlayer(player1, roll1Die)
		if newPlayer1.Score >= maxScore {
			player1Wins += count * roll1Universes
		} else {
			die2 := game.NewQuantumDie()
			for !die2.IsComplete() {
				roll2Die, roll2Universes := die2.GetDie()
				newPlayer2 := game.MovePlayer(player2, roll2Die)
				if newPlayer2.Score >= maxScore {
					player2Wins += count * roll1Universes * roll2Universes
				} else {
					newPlayer1Wins, newPlayer2Wins := processUniverse(newPlayer1, newPlayer2, count*roll1Universes*roll2Universes)
					player1Wins += newPlayer1Wins
					player2Wins += newPlayer2Wins
				}
				die2.IncrementState()
			}
		}
		die1.IncrementState()
	}

	return player1Wins, player2Wins
}

func main() {
	player1Position, player2Position, err := fr.ReadPositions("positions.txt")
	if err != nil {
		log.Fatal(err)
	}

	player1Wins, player2Wins := processUniverse(game.NewPlayer(player1Position), game.NewPlayer(player2Position), 1)

	fmt.Printf("Player 1 won %v times.\n", player1Wins)
	fmt.Printf("Player 2 won %v times.\n", player2Wins)
}
