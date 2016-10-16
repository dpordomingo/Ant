package geo

type Row []bool

type Map struct {
	Size
	World []Row
}

//NewMap returns a new Map
func NewMap() *Map {
	return &Map{}
}

//AddRow Ads a new row to the map
func (m *Map) AddRow(r Row) *Map {
	m.World = append(m.World, r)
	m.h++
	if rowLength := int32(len(r)); m.w < rowLength {
		m.w = rowLength
	}

	return m
}

//String returns a stringified version of the map
//TODO: dupe
func (m *Map) String() string {
	output := ""
	for r := range m.World {
		for c := range m.World[r] {
			if m.IsInside(Point{X: int32(c), Y: int32(r)}) {
				output += string([]byte{226, 150, 145})
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

//IsInside returns true if the passed point is in a valid location inside the map
func (m *Map) IsInside(point Point) bool {
	return int32(len(m.World[point.Y])) > point.X && m.World[point.Y][point.X]
}
