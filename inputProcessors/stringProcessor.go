package inputProcessors

import (
	"bufio"
	"os"

	"github.com/dpordomingo/learning-exercises/ant/actors"
	"github.com/dpordomingo/learning-exercises/ant/geo"
)

func Process(f *os.File) (
	geo.Map, geo.Point, actors.Ant) {

	var place geo.Map
	var mapTarget geo.Point
	var ant actors.Ant

	input := bufio.NewScanner(f)

	rowNumber := 0
	for input.Scan() {
		var rowMap geo.Row
		rowString := input.Text()
		for c := range rowString {
			char := rowString[c : c+1]
			if char == "X" {
				mapTarget.X = rowNumber
				mapTarget.Y = c
			}
			if char == "O" {
				ant.Origin = geo.Point{X: rowNumber, Y: c}
			}
			rowMap = append(rowMap, char != " ")
		}

		place.AddRow(rowMap)
		rowNumber++
		//fmt.Println(rowMap)
	}
	//fmt.Println(place)

	return place, mapTarget, ant
}
