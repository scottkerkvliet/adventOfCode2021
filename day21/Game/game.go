package game

const rollsPerTurn = 3

func MovePlayer(player Player, die Die) Player {
	roll := 0
	for i := 0; i < rollsPerTurn; i++ {
		roll += die.GetRoll()
	}
	newPosition := player.Position + (roll % 10)
	if newPosition > 10 {
		newPosition -= 10
	}
	player.Score += newPosition
	player.Position = newPosition

	return player
}
