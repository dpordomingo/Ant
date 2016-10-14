package generators

import "github.com/dpordomingo/learning-exercises/ant/geo"

type MapStringGenerator struct {
	size      geo.Size
	target    *geo.Point
	source    *geo.Point
	obstacles []obstacle
}

type obstacle struct {
	Point geo.Point
	geo.Area
}

func NewMapStringGenerator(size geo.Size) *MapStringGenerator {
	return &MapStringGenerator{
		size: size,
	}
}

//TODO: errores
// default values vs undefined values
func (g *MapStringGenerator) Generate() (*geo.Map, *geo.Point, *geo.Point, error) {
	generatedMap := g.getClearMap()
	g.addObstacles(generatedMap)

	if g.target == nil {
		g.target = GetPointInMap(generatedMap)
	}
	if g.source == nil {
		g.source = GetPointInMap(generatedMap)
	}

	return generatedMap, g.target, g.source, nil
}

func (g *MapStringGenerator) DefineSource(sourceX, sourceY int32) *MapStringGenerator {
	g.source = &geo.Point{sourceX, sourceY}
	return g
}

func (g *MapStringGenerator) DefineTarget(targetX, targetY int32) *MapStringGenerator {
	g.target = &geo.Point{targetX, targetY}
	return g
}

func (g *MapStringGenerator) DefineObstacle(count, sizeX, sizeY int32) *MapStringGenerator {
	for i := int32(1); i <= count; i++ {
		area := &geo.Area{g.size.W, g.size.H}
		size := &geo.Size{sizeX, sizeY}
		obstaclePoint, obstacleArea, _ := GetRandomIncludedArea(area, size)
		obstacle := obstacle{*obstaclePoint, obstacleArea}
		g.obstacles = append(g.obstacles, obstacle)
	}

	return g
}

func (g *MapStringGenerator) getClearMap() *geo.Map {
	generatedMap := geo.Map{}
	generatedMap.W = g.size.W
	generatedMap.H = g.size.H
	for j := int32(0); j < g.size.H; j++ {
		row := geo.Row{}
		for i := int32(0); i < g.size.W; i++ {
			row = append(row, true)
		}
		generatedMap.AddRow(row)
	}

	return &generatedMap
}

func (g *MapStringGenerator) addObstacles(givenMap *geo.Map) *geo.Map {
	for _, obstacle := range g.obstacles {
		point := obstacle.Point
		area := obstacle.Area
		for i := point.X; i < point.X+area.W; i++ {
			for j := point.Y; j < point.Y+area.H; j++ {
				givenMap.World[j][i] = false
			}
		}
	}
	return givenMap
}
