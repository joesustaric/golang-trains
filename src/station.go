package trains

type station struct {
	name        string
	connections map[*station]int
}

func (s station) AddConnection(con *station, distance int) {
	s.connections[con] = distance
}

func (s station) GetConnection(name string) station {
	for k := range s.connections {
		if k.name == name {
			return *k
		}
	}
	return station{"none", map[*station]int{}}
}

func (s station) GetDistanceTo(conn *station) int {
	return s.connections[conn]
}
