package geo

import "math"

//LocatedArea represents an area with a given coordinates for its top-left vertex
type LocatedArea struct {
	Point
	Area
}

//Point modelates a 2D point
type Point struct {
	X, Y int32
}

//Is returns true if passed coordinates corresponds with the referenced point
func (p *Point) Is(X, Y int32) bool {
	return p.X == X && p.Y == Y
}

//Equals returns true if passed point is equal the referenced point
func (p *Point) Equals(q Point) bool {
	return p.Is(q.X, q.Y)
}

//Area modelates a surface
type Area struct {
	w int32
	h int32
}

//NewArea returns a new Area
func NewArea(w, h int32) Area {
	return Area{w: int32(math.Abs(float64(w))), h: int32(math.Abs(float64(h)))}
}

//W returns the width of the Area
func (a Area) W() int32 {
	return a.w
}

//H returns the height of the Area
func (a Area) H() int32 {
	return a.h
}

//Size returns the Size of the Area
func (a Area) Size() Size {
	return Size(a)
}

//Size modelates a Size 2D dimmensions: (width X height)
type Size Area

//NewSize returns a new Size
func NewSize(w, h int32) Size {
	return Size{w: int32(math.Abs(float64(w))), h: int32(math.Abs(float64(h)))}
}

//W returns the width of the Size
func (s Size) W() int32 {
	return s.w
}

//H returns the height of the Size
func (s Size) H() int32 {
	return s.h
}

//BuildArea returns an Area with the referenced Size
func (s Size) BuildArea() Area {
	return Area(s)
}
