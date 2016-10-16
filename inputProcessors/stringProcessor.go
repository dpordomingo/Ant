package inputProcessors

import (
	"bufio"
	"io"

	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/literals"
)

//Process returns a Map, origin and destiny from STDin
func Process(f io.Reader) (*geo.Map, *geo.Point, *geo.Point, error) {

	var mapTarget *geo.Point
	var mapOrigin *geo.Point
	place := geo.NewMap()

	input := bufio.NewScanner(f)

	rowNumber := 0
	for input.Scan() {
		var rowMap geo.Row
		rowString := input.Text()
		i := 0
		for _, rune := range rowString {
			point := geo.Point{X: int32(i), Y: int32(rowNumber)}
			if string(rune) == "X" || string(rune) == literals.SYMBOL_ANT {
				mapTarget = &point
			}
			if string(rune) == "O" || string(rune) == literals.SYMBOL_POINT {
				mapOrigin = &point
			}
			rowMap = append(rowMap, string(rune) != " ")
			i++
		}

		place.AddRow(rowMap)
		rowNumber++
	}

	return place, mapOrigin, mapTarget, nil
}

type P struct {
	X int
}

func getP(empty bool) *P {
	if empty {
		return nil
	} else {
		return &P{2}
	}
}
