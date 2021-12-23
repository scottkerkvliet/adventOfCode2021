package snailnumber

import "fmt"

// ********************************* Interface ***********************************
type SnailNumber interface {
	GetMagnitude() int
	Print() string
	Copy() SnailNumber
	checkExplosions() bool
	checkSplits() bool
	explode(depth int) (SnailNumber, int, int)
	split() SnailNumber
	addLeft(value int) SnailNumber
	addRight(value int) SnailNumber
}

// ******************************* Constructors *********************************
func NewPair(left, right SnailNumber) SnailNumber {
	return &Pair{left, right}
}

func NewSingle(value int) SnailNumber {
	return Single(value)
}

func SplitValue(value int) SnailNumber {
	left := value / 2
	right := (value / 2) + (value % 2)
	return &Pair{Single(left), Single(right)}
}

func Add(left, right SnailNumber) SnailNumber {
	newNumber := &Pair{left, right}

	exploded, split := true, true
	for exploded || split {
		exploded = newNumber.checkExplosions()
		if !exploded {
			split = newNumber.checkSplits()
		}
	}

	return newNumber
}

// *********************************** Pair **************************************
type Pair struct {
	left, right SnailNumber
}

func (p *Pair) GetMagnitude() int {
	return (3 * p.left.GetMagnitude()) + (2 * p.right.GetMagnitude())
}

func (p *Pair) Print() string {
	return fmt.Sprintf("[%v,%v]", p.left.Print(), p.right.Print())
}

func (p *Pair) Copy() SnailNumber {
	left := p.left.Copy()
	right := p.right.Copy()
	return NewPair(left, right)
}

func (p *Pair) checkExplosions() bool {
	newPair, _, _ := p.explode(0)
	return newPair != nil
}

func (p *Pair) checkSplits() bool {
	return p.split() != nil
}

func (p *Pair) explode(depth int) (SnailNumber, int, int) {
	newLeft, l1, r1 := p.left.explode(depth + 1)
	if newLeft != nil {
		p.left = newLeft
		p.right = p.right.addLeft(r1)
		return p, l1, 0
	}

	newRight, l2, r2 := p.right.explode(depth + 1)
	if newRight != nil {
		p.right = newRight
		p.left = p.left.addRight(l2)
		return p, 0, r2
	}

	if depth >= 4 {
		return Single(0), p.left.GetMagnitude(), p.right.GetMagnitude()
	}

	return nil, 0, 0
}

func (p *Pair) split() SnailNumber {
	left := p.left.split()
	if left != nil {
		p.left = left
		return p
	}

	right := p.right.split()
	if right != nil {
		p.right = right
		return p
	}

	return nil
}

func (p *Pair) addLeft(value int) SnailNumber {
	if value != 0 {
		p.left = p.left.addLeft(value)
	}
	return p
}

func (p *Pair) addRight(value int) SnailNumber {
	if value != 0 {
		p.right = p.right.addRight(value)
	}
	return p
}

// ********************************** Single *************************************
type Single int

func (s Single) GetMagnitude() int {
	return int(s)
}

func (s Single) Print() string {
	return fmt.Sprint(s.GetMagnitude())
}

func (s Single) Copy() SnailNumber {
	return NewSingle(s.GetMagnitude())
}

func (s Single) checkExplosions() bool {
	return false
}

func (s Single) checkSplits() bool {
	return false
}

func (s Single) explode(depth int) (SnailNumber, int, int) {
	return nil, 0, 0
}

func (s Single) split() SnailNumber {
	if int(s) > 9 {
		return SplitValue(s.GetMagnitude())
	}

	return nil
}

func (s Single) addLeft(value int) SnailNumber {
	return s + Single(value)
}

func (s Single) addRight(value int) SnailNumber {
	return s + Single(value)
}
