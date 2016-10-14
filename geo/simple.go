package geo

type Point struct {
	X, Y int32
}

func (p *Point) Equals(X, Y int32) bool {
	return p.X == X && p.Y == Y
}

type Area struct {
	W int32
	H int32
}

type Size Area
