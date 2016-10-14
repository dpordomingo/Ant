package generators

import (
	"fmt"
	"math/rand"

	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/literals"
)

//RandomBoolean returns a random boolean that will be true with the passed probability
func RandomBoolean(trueFreq float64) bool {
	randNum := rand.Float64()
	return randNum < trueFreq
}

//GetRandomPoint returns a random Poiny{X,Y} where coordinates will be smaller than passed
// the point coordinates could be negative if passed coordinates are negative
func GetRandomPoint(maxX, maxY int32) *geo.Point {
	r := GetRandx()
	signX := int32(+1)
	signY := int32(+1)
	if maxX < 0 && RandomBoolean(.5) {
		signX = int32(-1)
	}
	if maxY < 0 && RandomBoolean(.5) {
		signY = int32(-1)
	}
	return &geo.Point{signX * r.Int31n(maxX), signY * r.Int31n(maxY)}
}

//GetRandomIncludedArea returns an area inside a parent one given the size of the new one,
// The returned area is defined by its top-left corner -randomly located- and defined size
func GetRandomIncludedArea(a *geo.Area, size *geo.Size) (*geo.Point, geo.Area, error) {
	if size.W > a.W || size.H > a.H {
		return &geo.Point{}, geo.Area{}, fmt.Errorf(literals.ERROR_AREA_OVERFLOW)
	}
	if size.W <= 0 || size.H <= 0 {
		return &geo.Point{}, geo.Area{}, fmt.Errorf(literals.ERROR_SIZE_NEGATIVE)
	}
	maxX := a.W - size.W
	maxY := a.H - size.H
	limitArea := &geo.Area{maxX + 1, maxY + 1}
	areaOrigin := GetRandomInsidePoint(limitArea)
	return areaOrigin, geo.Area{size.W, size.H}, nil
}

//GetRandomInsidePoint returns a point -randomly located- inside the passed area
func GetRandomInsidePoint(a *geo.Area) *geo.Point {
	return GetRandomPoint(a.W, a.H)
}

func GetPointInMap(m *geo.Map) *geo.Point {
	candidate := GetRandomInsidePoint(&m.Area)
	for !m.IsInside(candidate) {
		candidate = GetRandomInsidePoint(&m.Area)
	}

	return candidate
}
