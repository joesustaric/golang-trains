package trains

// AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
func getTestNetworkOfTrains() *Network {
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
	s := []*station{&a, &b, &c, &d, &e}
	network.AddNode(s...)

	return &network
}
