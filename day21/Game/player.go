package game

type Player struct {
	Score, Position int
}

func NewPlayer(position int) Player {
	return Player{0, position}
}
