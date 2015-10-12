package trains

import "sync"

// Station - todo desc
type Station struct {
	name        string
	connections map[*Station]int
	connMutex   *sync.RWMutex
}

// NewStation returns a new station
func NewStation(name string) *Station {
	return &Station{name: name, connections: make(map[*Station]int), connMutex: &sync.RWMutex{}}
}

// NewStations returns a list of new stations
func NewStations(names ...string) map[string]*Station {
	r := make(map[string]*Station)

	for _, name := range names {
		r[name] = NewStation(name)
	}

	return r
}

// Stationer - todo desc
type Stationer interface {
	AddConnection(con *Station, distance int)
	GetConnection(name string) Station
	GetDistanceTo(conn *Station) int
}

// AddConnection - todo
func (s *Station) AddConnection(con *Station, distance int) {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()
	s.connections[con] = distance
}

// GetConnection - todo
func (s *Station) GetConnection(name string) *Station {
	s.connMutex.RLock()
	defer s.connMutex.RUnlock()

	for c := range s.connections {
		if c.name == name {
			return c
		}
	}
	return &Station{}
}

// GetDistanceTo - todo
func (s *Station) GetDistanceTo(conn *Station) int {
	return s.connections[conn]
}
