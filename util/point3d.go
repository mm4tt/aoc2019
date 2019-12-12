package util

type Point3d struct {
	X, Y, Z int
}

func (p *Point3d) Add(p2 Point3d) Point3d {
	return Point3d{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

func (p *Point3d) Neg() Point3d {
	return Point3d{-p.X, -p.Y, -p.Z}
}

func (p *Point3d) Norm() int {
	return Abs(p.X) + Abs(p.Y) + Abs(p.Z)
}
