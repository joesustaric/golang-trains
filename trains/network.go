package trains

//Network Network of trains
type Network struct {
	nodes map[string]*Station
}

// Networker todo - description
type Networker interface {
	AddNode(stations ...*Station)
	GetNode(name string) (*Station, bool)
}

//NewNetwork returns a new network
func NewNetwork() Network {
	return Network{nodes: make(map[string]*Station)}
}

//AddNode adds a station to the network
func (n Network) AddNode(stations ...*Station) {
	for _, s := range stations {
		n.nodes[s.name] = s
	}
}

//GetNode will get the node
func (n Network) GetNode(name string) (*Station, bool) {
	node, ok := n.nodes[name]
	return node, ok
}
