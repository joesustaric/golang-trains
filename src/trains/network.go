package trains

//Network Network of trains
type Network struct {
	nodes map[string]*station
}

//NewNetwork returns a new network
func NewNetwork() Network {
	return Network{map[string]*station{}}
}

//AddNode adds a station to the network
func (n Network) AddNode(stations ...*station) {
	for _, s := range stations {
		n.nodes[s.name] = s
	}
}

//GetNode will get the node
func (n Network) GetNode(name string) (*station, bool) {
	node, ok := n.nodes[name]
	return node, ok
}
