package generators

import "github.com/dpordomingo/learning-exercises/ant/geo"

//Struct that can generate random maps
type MapGenerator struct {
	size      geo.Size
	area      geo.Area
	obstacles []geo.Size
}

//NewMapGenerator returns a MapGenerator
func NewMapGenerator(size geo.Size) *MapGenerator {
	return &MapGenerator{
		size: size,
		area: size.BuildArea(),
	}
}

//DefineObstacle adds the required number of obstacles of the given size to the generator
func (g *MapGenerator) DefineObstacles(count int32, size geo.Size) *MapGenerator {
	for i := int32(1); i <= count; i++ {
		g.obstacles = append(g.obstacles, size)
	}

	return g
}

//Generate returns the random Map
func (g *MapGenerator) Generate() (*geo.Map, error) {

	generatedMap := g.getClearMap()
	g.addObstacles(generatedMap)

	return generatedMap, nil
}

func (g *MapGenerator) getClearMap() *geo.Map {
	generatedMap := geo.NewMap()
	for j := int32(0); j < g.size.H(); j++ {
		row := geo.Row{}
		for i := int32(0); i < g.size.W(); i++ {
			row = append(row, true)
		}

		generatedMap.AddRow(row)
	}

	return generatedMap
}

func (g *MapGenerator) addObstacles(givenMap *geo.Map) *geo.Map {
	for _, obstacle := range g.getRandomObstacles() {
		for x := obstacle.X; x < obstacle.X+obstacle.W(); x++ {
			for y := obstacle.Y; y < obstacle.Y+obstacle.H(); y++ {
				if givenMap.IsInside(geo.Point{X: x, Y: y}) {
					givenMap.World[y][x] = false
				}
			}
		}
	}

	return givenMap
}

func (g *MapGenerator) getRandomObstacles() []geo.LocatedArea {
	var obstacles []geo.LocatedArea
	for i := range g.obstacles {
		obstacle, _ := GetRandomIncludedArea(g.area, g.obstacles[i])
		obstacles = append(obstacles, obstacle)
	}

	return obstacles
}
