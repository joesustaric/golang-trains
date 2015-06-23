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

func TestCalculateConnectionDistanceFromCurrent(t *testing.T) {
	Convey("Given a valid network and Trip and initalised NewShortestPath", t, func() {
		testNetwork, _ := getSimplerTestNetworkOfTrains()
		org, _ := testNetwork.GetNode("A")
		dest, _ := testNetwork.GetNode("B")
		testTrip := &trip{org, dest}
		spt := NewShortestPathTrip(testNetwork, testTrip)

		Convey("When we ask to calculate distance to connecting nodes form current", func() {
			spt.CalculateConnectionDistanceFromCurrent()

			Convey("It calcualte the correct distances", func() {
				c, _ := testNetwork.GetNode("C")
				So(spt.distanceToStation[dest], ShouldEqual, 1)
				So(spt.distanceToStation[c], ShouldEqual, 4)
			})
		})
	})
}

func TestVisitNextNode(t *testing.T) {
	Convey("Given a valid NewShortestPath obejct and one set of distance calculations", t, func() {
		testNetwork, _ := getSimplerTestNetworkOfTrains()
		org, _ := testNetwork.GetNode("A")
		dest, _ := testNetwork.GetNode("B")
		testTrip := &trip{org, dest}
		spt := NewShortestPathTrip(testNetwork, testTrip)
		spt.CalculateConnectionDistanceFromCurrent()

		Convey("When we ask to move to the next node", func() {
			spt.VisitNextNode()

			Convey("It sets the next unvisited connection with the lowest connecting distance as the current and marks it as visited once", func() {
				b, _ := testNetwork.GetNode("B")
				connVistited, _ := spt.visitedStationTimes[b]
				So(connVistited, ShouldEqual, 1)
				So(spt.currentNode, ShouldEqual, dest)
				So(spt.distanceToStation[b], ShouldEqual, 1)
			})
		})
	})
}
