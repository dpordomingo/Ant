package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dpordomingo/learning-exercises/ant/actors"
	"github.com/dpordomingo/learning-exercises/ant/generators"
	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/inputProcessors"
	"github.com/dpordomingo/learning-exercises/ant/ui"
)

var (
	mapName  = flag.String("mapname", "", "Map name (stored in './maps' folder)")
	useStdIn = flag.Bool("stdin", false, "Map name (stored in './maps' folder)")
)

func main() {
	flag.Parse()
	worldMap, worldTaget, antOrigin, err := getMapOriginTarget()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ant := actors.NewAnt(worldMap)
	ant.Init(*antOrigin, *worldTaget)

	ui.PrintRepresentation(worldMap, worldTaget, ant)
	ui.RunServer(worldMap, worldTaget, ant)
}

func getMapOriginTarget() (*geo.Map, *geo.Point, *geo.Point, error) {
	var worldMap *geo.Map
	var worldTaget *geo.Point
	var antOrigin *geo.Point
	var err error

	if *useStdIn || *mapName != "" {
		worldMap, worldTaget, antOrigin, err = getStringMap()
	} else {
		worldMap, err = getRandomMap()
	}

	if err != nil {
		return nil, nil, nil, err
	}

	setOriginTargetIfUndefined(worldMap, &antOrigin, &worldTaget)
	return worldMap, worldTaget, antOrigin, err
}

func getRandomMap() (*geo.Map, error) {
	mapGenerator := generators.NewMapGenerator(geo.NewSize(60, 20))
	mapGenerator.DefineObstacles(2, geo.NewSize(8, 17))
	mapGenerator.DefineObstacles(1, geo.NewSize(8, 8))
	mapGenerator.DefineObstacles(1, geo.NewSize(4, 4))
	mapGenerator.DefineObstacles(3, geo.NewSize(2, 3))
	mapGenerator.DefineObstacles(20, geo.NewSize(1, 1))

	return mapGenerator.Generate()
}

func getStringMap() (*geo.Map, *geo.Point, *geo.Point, error) {
	var (
		source io.Reader
		err    error
	)
	if *useStdIn {
		source = os.Stdin
	} else {
		source, err = os.Open(fmt.Sprintf("maps/%s.map", *mapName))
		if err != nil {
			return nil, nil, nil, err
		}
	}

	m, o, t, e := inputProcessors.Process(source)
	if *useStdIn {
		fmt.Println("[EOF]\n")
	}
	fmt.Println("Map:")
	return m, o, t, e
}

func setOriginTargetIfUndefined(m *geo.Map, origin **geo.Point, target **geo.Point) {
	if *origin == nil {
		p := generators.GetPointInMap(m)
		*origin = &p
	}

	if *target == nil {
		p := generators.GetPointInMap(m)
		*target = &p
	}
}
