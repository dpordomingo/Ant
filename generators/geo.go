package generators

import (
	"fmt"
	"math/rand"

	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/literals"
)

//RandomBoolean returns a random boolean that will be true with the passed probability (0..1)
func RandomBoolean(trueFreq float64) bool {
	randNum := rand.Float64()
	return randNum < trueFreq
}

//GetRandomPoint returns a random Poiny{X,Y} where coordinates will be smaller than passed
// the point coordinates could be negative if passed coordinates are negative
func GetRandomPoint(maxX, maxY int32) geo.Point {
	r := GetRandx()
	signX := int32(+1)
	signY := int32(+1)

	if maxX < 0 && RandomBoolean(.5) {
		signX = int32(-1)
	}

	if maxY < 0 && RandomBoolean(.5) {
		signY = int32(-1)
	}

	return geo.Point{X: signX * r.Int31n(maxX), Y: signY * r.Int31n(maxY)}
}

//GetRandomInsidePoint returns a point -randomly located- inside the passed area
func GetRandomInsidePoint(a geo.Area) geo.Point {
	return GetRandomPoint(a.W(), a.H())
}

//GetRandomIncludedArea returns an area inside a parent one given the size of the new one,
// The returned area is defined by its top-left corner -randomly located- and defined size
func GetRandomIncludedArea(area geo.Area, size geo.Size) (geo.LocatedArea, error) {
	if size.W() > area.W() || size.H() > area.H() {
		return geo.LocatedArea{}, fmt.Errorf(literals.ERROR_AREA_OVERFLOW)
	}

	maxX := area.W() - size.W() + 1
	maxY := area.H() - size.H() + 1
	limitArea := geo.NewArea(maxX, maxY)
	areaOrigin := GetRandomInsidePoint(limitArea)
	return geo.LocatedArea{areaOrigin, size.BuildArea()}, nil
}

//GetPointInMap returns a Point in a valid position inside a Map
func GetPointInMap(m *geo.Map) geo.Point {
	mapArea := m.Size.BuildArea()
	candidate := GetRandomInsidePoint(mapArea)
	for !m.IsInside(candidate) {
		candidate = GetRandomInsidePoint(mapArea)
	}

	return candidate
}
