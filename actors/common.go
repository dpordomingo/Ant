package actors

import "github.com/dpordomingo/learning-exercises/ant/geo"

//Rover modelates an actor that can walk over a map and look for a target
type Rover interface {
	//Init initializes the Ant searching for certain destiny from certain origin
	Init(geo.Point, geo.Point)
	//Position returns the position of the Rover
	Position() geo.Point
}
