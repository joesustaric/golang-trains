package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewShortestPathTrip(t *testing.T) {

	Convey("Given a valid network and Trip", t, func() {
		testNetwork, _ := getSimplerTestNetworkOfTrains()
		org, _ := testNetwork.GetNode("A")
		dest, _ := testNetwork.GetNode("B")
		testTrip := &trip{org, dest}
		Convey("When we ask for a new Shortest Path Trip object", func() {
			spt := NewShortestPathTrip(testNetwork, testTrip)

			Convey("It initalises correctly", func() {
				So(spt.myNetwork, ShouldEqual, testNetwork)
				So(spt.originalTrip, ShouldEqual, testTrip)
				So(spt.visitedStationTimes, ShouldNotBeNil)
				So(spt.distanceToStation, ShouldNotBeNil)
			})
			Convey("It marks the origin as visited and the current node as the origin", func() {
				orgVisited, ok := spt.visitedStationTimes[org]
				So(ok, ShouldBeTrue)
				So(orgVisited, ShouldEqual, 1)
				So(spt.currentNode, ShouldEqual, org)
			})
		})
	})

}

func TestCanGetToDestination(t *testing.T) {

}
