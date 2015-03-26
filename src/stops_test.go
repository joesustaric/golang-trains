package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTripsBetween(t *testing.T) {

	Convey("Given a network of stations", t, func() {
		network := NewNetwork()
		network.AddNode(makeStations()...)
		Convey("When we ask for trips between with stop constraint of 3", func() {
			t := tripHopsQuery{origin: "C", destination: "C", maxHops: 3}
			result := GetNumberOfRoutes(network, t)
			Convey("It returns an array with all the correct routes", func() {
				So(result, ShouldEqual, 2)
			})
		})
	})
}

func makeStations() []*station {
	b := station{"B", map[*station]int{}}
	c := station{"C", map[*station]int{}}
	d := station{"D", map[*station]int{}}
	e := station{"E", map[*station]int{}}
	result := []*station{&b, &c, &d, &e}
	c.AddConnection(&d, 8)
	c.AddConnection(&e, 2)
	d.AddConnection(&e, 6)
	d.AddConnection(&c, 8)
	e.AddConnection(&b, 3)
	b.AddConnection(&c, 4)
	return result
}
