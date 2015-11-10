package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateShortestDistance(t *testing.T) {
	Convey("When we have a network of stations and a possible trip A-C", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, dest := n.GetStation("A"), n.GetStation("C")
		t := trip{org, dest}
		shortestPathTrip, _ := NewShortestPathTrip(n, &t)

		Convey("When we ask it to calculate the shortest distance", func() {
			CalculateShortestTrip(shortestPathTrip)

			Convey("It calculates the correct shortest path", func() {
				So(shortestPathTrip.visitedSet[dest], ShouldEqual, 9)
			})
		})
	})

	Convey("When we have a network of stations and a possible trip B-B", t, func() {
		n, _ := getTestNetworkOfTrains()
		org := n.GetStation("B")
		t := trip{org, org}
		shortestPathTrip, _ := NewShortestPathTrip(n, &t)

		Convey("When we ask it to calculate the shortest distance", func() {
			CalculateShortestTrip(shortestPathTrip)

			Convey("It calculates the correct shortest path", func() {
				So(shortestPathTrip.visitedSet[org], ShouldEqual, 9)
			})
		})
	})
}
