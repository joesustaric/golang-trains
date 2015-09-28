package trains

// AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
func getTestNetworkOfTrains() (*Network, []*Station) {
	network := NewNetwork()

	a := Station{"A", map[*Station]int{}}
	b := Station{"B", map[*Station]int{}}
	c := Station{"C", map[*Station]int{}}
	d := Station{"D", map[*Station]int{}}
	e := Station{"E", map[*Station]int{}}

	a.AddConnection(&b, 5)
	b.AddConnection(&c, 4)
	c.AddConnection(&d, 8)
	d.AddConnection(&c, 8)
	d.AddConnection(&e, 6)
	a.AddConnection(&d, 5)
	c.AddConnection(&e, 2)
	e.AddConnection(&b, 3)
	a.AddConnection(&e, 7)
	stations := []*Station{&a, &b, &c, &d, &e}
	network.AddNode(stations...)

	return &network, stations
}

// AB3, BC5, BD3, BE1, EC1, CF7
func getSimpleTestNetworkOfTrains() (*Network, []*Station) {
	network := NewNetwork()

	a := Station{"A", map[*Station]int{}}
	b := Station{"B", map[*Station]int{}}
	c := Station{"C", map[*Station]int{}}
	d := Station{"D", map[*Station]int{}}
	e := Station{"E", map[*Station]int{}}
	f := Station{"F", map[*Station]int{}}

	a.AddConnection(&b, 3)
	b.AddConnection(&c, 5)
	b.AddConnection(&d, 3)
	b.AddConnection(&e, 1)
	e.AddConnection(&c, 1)
	c.AddConnection(&f, 7)

	Stations := []*Station{&a, &b, &c, &d, &e, &f}
	network.AddNode(Stations...)

	return &network, Stations
}

// AB1, AC4, BC2, DB5, CA3
func getSimplerTestNetworkOfTrains() (*Network, []*Station) {
	network := NewNetwork()

	a := Station{"A", map[*Station]int{}}
	b := Station{"B", map[*Station]int{}}
	c := Station{"C", map[*Station]int{}}
	d := Station{"D", map[*Station]int{}}

	a.AddConnection(&b, 1)
	a.AddConnection(&c, 4)
	b.AddConnection(&c, 2)
	b.AddConnection(&d, 5)
	c.AddConnection(&a, 3)

	Stations := []*Station{&a, &b, &c, &d}
	network.AddNode(Stations...)

	return &network, Stations
}
