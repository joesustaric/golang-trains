package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistanceCalc(t *testing.T) {

	Convey("Given a Network with 2 connecting stations", t, func() {
		a := station{"A", map[*station]int{}}
		b := station{"B", map[*station]int{}}
		a.AddConnection(&b, 5)
		query := []string{"A", "B"}
		Network := NewNetwork()
		Network.AddNode(&a, &b)
		Convey("When we ask for the distance from A-B", func() {
			result := TotalDistance(Network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 5)
			})
		})
	})

	Convey("Given a Network with 3 connecting stations", t, func() {
		stations := makeDistanceTestStations()
		Network := NewNetwork()
		Network.AddNode(stations...)
		Convey("When we ask for the distance from A-B-C", func() {
			query := []string{"A", "B", "C"}
			result := TotalDistance(Network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 9)
			})
		})

		Convey("When we ask for the distance to a station it does not connect to", func() {
			query := []string{"A", "B", "D"}
			result := TotalDistance(Network, query)
			Convey("It returns 0", func() {
				So(result, ShouldEqual, 0)
			})
		})
	})
}

func makeDistanceTestStations() []*station {
	a := station{"A", map[*station]int{}}
	b := station{"B", map[*station]int{}}
	c := station{"C", map[*station]int{}}
	a.AddConnection(&b, 5)
	b.AddConnection(&c, 4)
	result := []*station{&a, &b, &c}
	return result
}
