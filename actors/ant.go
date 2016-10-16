package actors

import (
	"github.com/dpordomingo/learning-exercises/ant/geo"
)

//Ant modelates a Rover
type Ant struct {
	ID         int
	world      *geo.Map
	origin     geo.Point
	destiny    geo.Point
	currentPos geo.Point
	path       []geo.Point
}

//NewAnt returns a new Ant
func NewAnt(world *geo.Map) Rover {
	return &Ant{
		world: world,
	}
}

//Init initializes the Ant searching for certain destiny from certain origin
func (a *Ant) Init(origin geo.Point, destiny geo.Point) {
	a.origin = origin
	a.destiny = destiny
	a.currentPos = a.origin
	a.path = append(a.path, a.origin)
	a.walk()
}

//Position returns the position of the Rover
func (a *Ant) Position() geo.Point {
	return a.currentPos
}

func (a *Ant) walk() {
	//TODO: code here
}
