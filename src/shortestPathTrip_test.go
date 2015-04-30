package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewShortestPathTrip(t *testing.T) {

	Convey("Given a network and origin and destination trip", t, func() {

		n := getTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("B")
		t := trip{org, dest}
		Convey("When we ask for a new ShortestPathTrip Object", func() {
			shortestPathTrip := NewShortestPathTrip(n, &t)
			Convey("It returns a correctly initalised object", func() {
				visitedStations := getVisitedStations(n)
				So(shortestPathTrip.network, ShouldEqual, n)
				So(shortestPathTrip.originalTrip, ShouldEqual, &t)
				So(shortestPathTrip.visitedStations, ShouldResemble, visitedStations)
			})
		})

	})
}

func getVisitedStations(network *Network) map[*station]bool {
	r := make(map[*station]bool)

	for _, s := range network.nodes {
		r[s] = false
	}

	return r
}
