package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateShortestDistance(t *testing.T) {
	SkipConvey("When we have a network of stations and a possible trip A-C", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("C")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)

		Convey("When we ask it to calculate the shortest distance", func() {
			CalculateShortestTrip(shortestPathTrip)

			Convey("It calculates the correct shortest path", func() {
				So(shortestPathTrip.distanceToStation[dest], ShouldEqual, 9)
			})
		})
	})

	Convey("When we have a network of stations and a possible trip B-B", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("B")
		dest, _ := n.GetNode("B")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)

		Convey("When we ask it to calculate the shortest distance", func() {
			CalculateShortestTrip(shortestPathTrip)

			Convey("It calculates the correct shortest path", func() {
				So(shortestPathTrip.distanceToStation[dest], ShouldEqual, 9)
			})
		})
	})
}
