package game

type DeterministicDie struct {
	lastRoll, count int
}

func NewDeterministicDie() *DeterministicDie {
	return &DeterministicDie{0, 0}
}

func (d *DeterministicDie) GetRoll() int {
	d.count++
	d.lastRoll++
	if d.lastRoll > 100 {
		d.lastRoll = 1
	}
	return d.lastRoll
}

func (d *DeterministicDie) GetCount() int {
	return d.count
}
