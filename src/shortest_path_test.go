package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetShortestRoute(t *testing.T) {

	Convey("Given a network of stations", t, func() {
		network := NewNetwork()
		network.AddNode(makeTestStationsA()...)
		Convey("When we ask for the shortest path from A to E", func() {
			result := GetShortestRoute(network, "A", "E")
			Convey("It returns the shortest path", func() {
				So(result, ShouldEqual, 5)
			})
		})
	})

	// Convey("Given a network of stations", t, func() {
	// 	network := NewNetwork()
	// 	network.AddNode(makeTestStationsB()...)
	// 	Convey("When we ask for the shortest path from A to C", func() {
	// 		result := GetShortestRoute(network, "A", "C")
	// 		Convey("It returns the shortest path", func() {
	// 			So(result, ShouldEqual, 9)
	// 		})
	// 	})
	// })

}

func makeTestStationsA() []*station {
	a := station{"A", map[*station]int{}}
	b := station{"B", map[*station]int{}}
	c := station{"C", map[*station]int{}}
	d := station{"D", map[*station]int{}}
	e := station{"E", map[*station]int{}}
	result := []*station{&a, &b, &c, &d, &e}
	a.AddConnection(&b, 1)
	a.AddConnection(&c, 1)
	a.AddConnection(&d, 1)
	b.AddConnection(&e, 5)
	b.AddConnection(&c, 2)
	c.AddConnection(&e, 4)
	c.AddConnection(&d, 2)
	d.AddConnection(&e, 6)
	return result
}

// func makeTestStationsB() []*station {
// 	a := station{"A", map[*station]int{}}
// 	b := station{"B", map[*station]int{}}
// 	c := station{"C", map[*station]int{}}
// 	d := station{"D", map[*station]int{}}
// 	e := station{"E", map[*station]int{}}
// 	result := []*station{&a, &b, &c, &d, &e}
// 	a.AddConnection(&b, 5)
// 	b.AddConnection(&c, 4)
// 	c.AddConnection(&d, 8)
// 	d.AddConnection(&c, 8)
// 	d.AddConnection(&e, 6)
// 	a.AddConnection(&d, 5)
// 	c.AddConnection(&e, 2)
// 	e.AddConnection(&b, 3)
// 	a.AddConnection(&e, 7)
// 	return result
// }
