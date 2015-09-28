package trains

// Station - todo desc
type Station struct {
	name        string
	connections map[*Station]int
}

// Stationer - todo desc
type Stationer interface {
	AddConnection(con *Station, distance int)
	GetConnection(name string) Station
	GetDistanceTo(conn *Station) int
}

// AddConnection - todo
func (s Station) AddConnection(con *Station, distance int) {
	s.connections[con] = distance
}

// GetConnection - todo
func (s Station) GetConnection(name string) Station {
	for k := range s.connections {
		if k.name == name {
			return *k
		}
	}
	return Station{}
}

// GetDistanceTo - todo
func (s Station) GetDistanceTo(conn *Station) int {
	return s.connections[conn]
}
