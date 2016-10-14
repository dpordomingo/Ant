package main

import (
	"fmt"

	"github.com/dpordomingo/learning-exercises/ant/generators"
	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/ui"
)

func main() {
	var worldMap *geo.Map
	var worldTaget *geo.Point
	var origin *geo.Point

	mapGenerator := generators.NewMapStringGenerator(geo.Size{60, 20})

	mapGenerator.DefineObstacle(2, 8, 17)
	mapGenerator.DefineObstacle(1, 8, 8)
	mapGenerator.DefineObstacle(1, 4, 4)
	mapGenerator.DefineObstacle(3, 2, 3)
	mapGenerator.DefineObstacle(20, 1, 1)

	worldMap, worldTaget, origin, _ = mapGenerator.Generate()

	fmt.Println(ui.GetRepresentation(worldMap, worldTaget, origin))
}
