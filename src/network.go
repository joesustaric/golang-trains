package trains

//Network Network of trains
type Network struct {
	nodes map[string]*station
}

//NewNetwork does what it says
func NewNetwork() Network {
	return Network{map[string]*station{}}
}

//AddNode adds nodes
func (n Network) AddNode(stations ...*station) {
	for _, s := range stations {
		n.nodes[s.name] = s
	}
}

//GetNode gets Nodes
func (n Network) GetNode(name string) (*station, bool) {
	node, ok := n.nodes[name]
	return node, ok
}
