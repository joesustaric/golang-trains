package trains

type network struct {
	nodes map[string]*station
}

func (n network) AddNode(s *station) {
	n.nodes[s.name] = s
}

func (n network) GetNode(name string) *station {
	return n.nodes[name]
}