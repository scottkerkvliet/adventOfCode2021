package cuboid

type Instruction struct {
	On bool
	Cuboid *Cuboid
}

func NewInstruction(x1, x2, y1, y2, z1, z2 int, on bool) *Instruction {
	return &Instruction{on, NewCuboid(x1, x2, y1, y2, z1, z2)}
}

func (i *Instruction) Print() string {
	p := ""

	if i.On {
		p += "On "
	} else {
		p += "Off "
	}

	if i.Cuboid != nil {
		p += i.Cuboid.Print()
	} else {
		p += "(nil)"
	}

	return p
}
