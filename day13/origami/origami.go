package origami

type Fold struct {
	At int
	X bool
}

type Point struct {
	X, Y int
}

type Origami struct {
	Points []Point
}

func(o *Origami) Fold(fold Fold) {
	if fold.X {
		o.FoldX(fold.At)
	} else {
		o.FoldY(fold.At)
	}
}

func(o *Origami) FoldX(x int) {
	var newPoints []Point
	pointSet := map[Point]bool{}

	for _, point := range o.Points {
		if point.X > x {
			point.X = x - (point.X - x)
		}

		if _, exists := pointSet[point]; exists {
			continue
		}

		newPoints = append(newPoints, point)
		pointSet[point] = true
	}

	o.Points = newPoints
}

func(o *Origami) FoldY(y int) {
	var newPoints []Point
	pointSet := map[Point]bool{}

	for _, point := range o.Points {
		if point.Y > y {
			point.Y = y - (point.Y - y)
		}

		if _, exists := pointSet[point]; exists {
			continue
		}

		newPoints = append(newPoints, point)
		pointSet[point] = true
	}

	o.Points = newPoints
}
