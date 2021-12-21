package game

type NumberDie struct {
	num int
}

func NewNumberDie(num int) *NumberDie {
	return &NumberDie{num}
}

func (d *NumberDie) GetRoll() int {
	roll := d.num
	d.num = 0
	return roll
}
