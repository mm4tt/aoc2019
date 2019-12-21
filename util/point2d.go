package util

type Point2D struct {
	X, Y int
}

func (p *Point2D) Add(p2 Point2D) Point2D {
	return Point2D{p.X + p2.X, p.Y + p2.Y}
}

func (p *Point2D) RotateLeft() Point2D {
	return Point2D{p.Y, -p.X}
}

func (p *Point2D) RotateRight() Point2D {
	return Point2D{-p.Y, p.X}
}

func (p *Point2D) Dist(p2 Point2D) int {
	return Abs(p.X - p2.X) + Abs(p.Y - p2.Y)
}
