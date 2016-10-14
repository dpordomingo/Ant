package geo

type Row []bool

type Map struct {
	Area
	World []Row
}

func (m *Map) AddRow(r Row) *Map {
	m.World = append(m.World, r)
	return m
}

//TODO: dupe
func (m *Map) String() string {
	output := ""
	for r := range m.World {
		for c := range m.World[r] {
			if m.World[r][c] {
				output += string([]byte{226, 150, 145})
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}
func (m *Map) IsInside(point *Point) bool {
	return m.World[point.Y][point.X]
}
