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
			Convey("It contains the correct refrence to network", func() {
				So(shortestPathTrip.network, ShouldEqual, n)
			})
			Convey("It contains the correct refrence to the original trip", func() {
				So(shortestPathTrip.originalTrip, ShouldEqual, &t)
			})
			Convey("It initialises the visited stations correctly", func() {
				visitedStations := getVisitedStations(n, &t)
				So(shortestPathTrip.visitedStations, ShouldResemble, visitedStations)
			})
			Convey("It initialises the distance to all nodes as infinity(9999) except origin", func() {
				distanceToStationFromOrigin := getStationsDistanceFromOrigin(shortestPathTrip.visitedStations, &t)
				So(shortestPathTrip.distanceToStation, ShouldResemble, distanceToStationFromOrigin)
			})
		})

	})
}

func TestNextStationToVisit(t *testing.T) {

	Convey("Given a network and origin and destination trip", t, func() {

		n := getTestNetworkOfTrains()
		Convey("When we ask for a new NextStationToVisit when org and dest are different", func() {
			org, _ := n.GetNode("A")
			dest, _ := n.GetNode("B")
			t := trip{org, dest}
			shortestPathTrip := NewShortestPathTrip(n, &t)
			nextStation := shortestPathTrip.GetNextStation()
			Convey("It returns the correct next station ", func() {
				So(nextStation.name, ShouldEqual, "B")
			})

		})

		// Convey("When we ask for a new NextStationToVisit when org and dest are the same", func() {
		// 	org, _ := n.GetNode("B")
		// 	dest, _ := n.GetNode("B")
		// 	t := trip{org, dest}
		// 	shortestPathTrip := NewShortestPathTrip(n, &t)
		// 	Convey("It returns the correct next station ", func() {
		// 		// assert default next station
		// 	})
		//
		// })

	})

}

//these setup the state we expect in the test.
func getVisitedStations(network *Network, t *trip) map[*station]bool {
	r := make(map[*station]bool)
	// for _, s := range network.nodes {
	// 	r[s] = false
	// }
	r[t.from] = true
	return r
}

func getStationsDistanceFromOrigin(v map[*station]bool, t *trip) map[*station]int {
	r := make(map[*station]int)
	for s := range v {
		r[s] = 99999
	}
	r[t.from] = 0
	return r
}
