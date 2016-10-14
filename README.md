# Ant rover

Playing with Go

## Maps generated randomly:
Example:

	mapGenerator := generators.NewMapStringGenerator(geo.Size{60, 20})
	mapGenerator.DefineObstacle(2, 8, 17)
	mapGenerator.DefineObstacle(1, 8, 8)
	mapGenerator.DefineObstacle(1, 4, 4)
	mapGenerator.DefineObstacle(3, 2, 3)
	mapGenerator.DefineObstacle(20, 1, 1)
	worldMap, worldTaget, origin, _ = mapGenerator.Generate()
	fmt.Println(ui.GetRepresentation(worldMap, worldTaget, origin))

	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ░░░░ ░░░░░░░░░░░░░░░░
	░░░░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  ░░░░░░░░░░  ░░░
	░░░░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  ░░░░░░░ ░░  ░░░
	░░░░░░                 ░░░░░░░░░░░░░░░░░ ░░  ░░░░        ░░░
	░░░░░░        ░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░        ░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░         ░        ░░░░░░░ ░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░        ░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░        ░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░  ░░░░░░░        ░░░░░░░░░░░░░░ ░░░░░░░░░░░        ░░░
	░░░░░░░░░░░░░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░  ░░░░░░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░  ░░░░░░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░  ░░░░░░░░        ░░░░░░░░░░░░░░░░░░ ░░░░░░░        ░░░
	░░░░░░░░░░░░░          ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░ ░░░░░░░░░          ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░░🐜░░░░░          ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
	░░░░░░░░░ ░░░          ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░ 
	░░░░░░░░░░░░░░░        ░░░░░░░░ ░⭙░░░░░░░░░░░░░░░        ░░░
	░░░░░░░░░░░ ░░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░        ░░░
