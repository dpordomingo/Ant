package ui

import (
	"fmt"

	"github.com/dpordomingo/learning-exercises/ant/actors"
	"github.com/dpordomingo/learning-exercises/ant/geo"
	"github.com/dpordomingo/learning-exercises/ant/literals"
)

//GetRepresentation returns a stringified representation of a Map, target and Rover
func GetRepresentation(m *geo.Map, target *geo.Point, rover actors.Rover) string {
	output := ""
	for r := range m.World {
		for c := range m.World[r] {
			cursorPoint := geo.Point{X: int32(c), Y: int32(r)}
			if target != nil && cursorPoint.Equals(*target) {
				output += literals.SYMBOL_POINT
			} else if rover != nil && cursorPoint.Equals(rover.Position()) {
				output += literals.SYMBOL_ANT
			} else if m.IsInside(geo.Point{X: int32(c), Y: int32(r)}) {
				output += literals.SYMBOL_GRASS_LIGHT
			} else {
				output += literals.SYMBOL_GRASS_BLANK
			}
		}

		output += "\n"
	}

	return output
}

//PrintRepresentation writes over the STDoutput the stringified representation of a Map, target and Rover
func PrintRepresentation(m *geo.Map, target *geo.Point, rover actors.Rover) {
	fmt.Println(GetRepresentation(m, target, rover))
	fmt.Printf("Map size: %d,%d\n", m.W(), m.H())

	if rover != nil {
		fmt.Printf("Rover: %d,%d\n", rover.Position().X, rover.Position().Y)
	}

	if target != nil {
		fmt.Printf("Target: %d,%d\n", target.X, target.Y)
	}
	fmt.Println("\n")
}
