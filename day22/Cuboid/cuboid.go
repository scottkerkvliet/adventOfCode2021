package cuboid

import "fmt"

type Cuboid struct {
	X1, X2, Y1, Y2, Z1, Z2 int
}

func NewCuboid(x1, x2, y1, y2, z1, z2 int) *Cuboid {
	return &Cuboid{x1, x2, y1, y2, z1, z2}
}

func (c *Cuboid) GetVolume() int {
	return (c.X2-c.X1+1) * (c.Y2-c.Y1+1) * (c.Z2-c.Z1+1)
}

func (c *Cuboid) Print() string {
	return fmt.Sprintf("(%v..%v),(%v..%v),(%v..%v) (volume: %v)", c.X1, c.X2, c.Y1, c.Y2, c.Z1, c.Z2, c.GetVolume())
}

func (c *Cuboid) Intersects(nc *Cuboid) bool {
	if c.X1 > nc.X2 || c.X2 < nc.X1 || c.Y1 > nc.Y2 || c.Y2 < nc.Y1 || c.Z1 > nc.Z2 || c.Z2 < nc.Z1 {
		return false
	}
	return true
}

func RemoveFromCuboid(base, top *Cuboid) []*Cuboid {
	var newCuboids []*Cuboid

	if !base.Intersects(top) {
		return []*Cuboid{base}
	}

	if top.X1 <= base.X2 && top.X1 > base.X1 {
		newCuboids = append(newCuboids, NewCuboid(base.X1, top.X1-1, base.Y1, base.Y2, base.Z1, base.Z2))
		base.X1 = top.X1
	}
	if top.X2 >= base.X1 && top.X2 < base.X2 {
		newCuboids = append(newCuboids, NewCuboid(top.X2+1, base.X2, base.Y1, base.Y2, base.Z1, base.Z2))
		base.X2 = top.X2
	}

	if top.Y1 <= base.Y2 && top.Y1 > base.Y1 {
		newCuboids = append(newCuboids, NewCuboid(base.X1, base.X2, base.Y1, top.Y1-1, base.Z1, base.Z2))
		base.Y1 = top.Y1
	}
	if top.Y2 >= base.Y1 && top.Y2 < base.Y2 {
		newCuboids = append(newCuboids, NewCuboid(base.X1, base.X2, top.Y2+1, base.Y2, base.Z1, base.Z2))
		base.Y2 = top.Y2
	}

	if top.Z1 <= base.Z2 && top.Z1 > base.Z1 {
		newCuboids = append(newCuboids, NewCuboid(base.X1, base.X2, base.Y1, base.Y2, base.Z1, top.Z1-1))
		base.Z1 = top.Z1
	}
	if top.Z2 >= base.Z1 && top.Z2 < base.Z2 {
		newCuboids = append(newCuboids, NewCuboid(base.X1, base.X2, base.Y1, base.Y2, top.Z2+1, base.Z2))
		base.Z2 = top.Z2
	}

	return newCuboids
}
