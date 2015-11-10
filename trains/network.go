package trains

//Network Network of trains
type Network struct {
	stations map[string]*Station
}

// Networker todo - description
type Networker interface {
	AddNode(stations ...*Station)
	GetNode(name string) *Station
}

//NewNetwork returns a new network
func NewNetwork() Network {
	return Network{stations: make(map[string]*Station)}
}

//AddStation adds a station to the network
func (n Network) AddStation(stations ...*Station) {
	for _, s := range stations {
		n.stations[s.name] = s
	}
}

//GetStation will return the Station if it exists.
//If not you get a nil.
func (n Network) GetStation(name string) *Station {
	if station, ok := n.stations[name]; ok {
		return station
	} else {
		return nil
	}
}
