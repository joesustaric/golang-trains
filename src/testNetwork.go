package trains

// AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
func getTestNetworkOfTrains() (*Network, []*station) {
	network := NewNetwork()

	a := station{"A", map[*station]int{}}
	b := station{"B", map[*station]int{}}
	c := station{"C", map[*station]int{}}
	d := station{"D", map[*station]int{}}
	e := station{"E", map[*station]int{}}

	a.AddConnection(&b, 5)
	b.AddConnection(&c, 4)
	c.AddConnection(&d, 8)
	d.AddConnection(&c, 8)
	d.AddConnection(&e, 6)
	a.AddConnection(&d, 5)
	c.AddConnection(&e, 2)
	e.AddConnection(&b, 3)
	a.AddConnection(&e, 7)
	stations := []*station{&a, &b, &c, &d, &e}
	network.AddNode(stations...)

	return &network, stations
}

// AB3, BC5, BD3, BE1, EC1, CF7
func getSimpleTestNetworkOfTrains() (*Network, []*station) {
	network := NewNetwork()

	a := station{"A", map[*station]int{}}
	b := station{"B", map[*station]int{}}
	c := station{"C", map[*station]int{}}
	d := station{"D", map[*station]int{}}
	e := station{"E", map[*station]int{}}
	f := station{"F", map[*station]int{}}

	a.AddConnection(&b, 3)
	b.AddConnection(&c, 5)
	b.AddConnection(&d, 3)
	b.AddConnection(&e, 1)
	e.AddConnection(&c, 1)
	c.AddConnection(&f, 7)

	stations := []*station{&a, &b, &c, &d, &e, &f}
	network.AddNode(stations...)

	return &network, stations
}
