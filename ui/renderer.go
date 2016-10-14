package ui

import "github.com/dpordomingo/learning-exercises/ant/geo"

var (
	SYMBOL_GRASS_BLANK  string = " "
	SYMBOL_POINT        string = string([]byte{226, 173, 153})
	SYMBOL_QUEEN        string = string([]byte{226, 153, 148})
	SYMBOL_ANT          string = string([]byte{240, 159, 144, 156})
	SYMBOL_GRASS_LIGHT  string = string([]byte{226, 150, 145})
	SYMBOL_GRASS_MEDIUM string = string([]byte{226, 150, 146})
	SYMBOL_GRASS_DARK   string = string([]byte{226, 150, 147})
)

//http://unicode-table.com/en/#2654
func GetRepresentation(m *geo.Map, target *geo.Point, ant *geo.Point) string {
	output := ""
	for r := range m.World {
		for c := range m.World[r] {
			if target.Equals(int32(c), int32(r)) {
				output += SYMBOL_POINT
			} else if ant.Equals(int32(c), int32(r)) {
				output += SYMBOL_ANT
			} else if m.World[r][c] {
				output += SYMBOL_GRASS_LIGHT
			} else {
				output += SYMBOL_GRASS_BLANK
			}
		}
		output += "\n"
	}
	return output
}
