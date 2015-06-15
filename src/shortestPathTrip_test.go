package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewShortestPathTrip(t *testing.T) {

	Convey("Given a network and origin and destination trip", t, func() {

		n, sts := getTestNetworkOfTrains()
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
			Convey("It contains the correct current node", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, org)
			})
			Convey("It initialises the distance to all nodes as infinity(9999) except origin", func() {
				distanceToStationFromOrigin := getStationsDistanceFromOrigin(sts, &t)
				So(shortestPathTrip.distanceToStation, ShouldResemble, distanceToStationFromOrigin)
			})
		})

	})
}

func TestVisitNextStation(t *testing.T) {

	Convey("Given a network and origin and destination trip", t, func() {
		n, _ := getTestNetworkOfTrains()
		Convey("When we ask for a new NextStationToVisit when org and dest are different", func() {
			org, _ := n.GetNode("A")
			dest, _ := n.GetNode("B")
			t := trip{org, dest}
			shortestPathTrip := NewShortestPathTrip(n, &t)
			shortestPathTrip.VisitNextStation()
			Convey("It returns the correct next station ", func() {
				So(shortestPathTrip.currentNode.name, ShouldBeIn, []string{"B", "D"})
			})
		})

		Convey("When we ask for a new NextStationToVisit when org and dest are the same", func() {
			org, _ := n.GetNode("B")
			dest, _ := n.GetNode("B")
			t := trip{org, dest}
			shortestPathTrip := NewShortestPathTrip(n, &t)
			shortestPathTrip.VisitNextStation()
			Convey("It returns the correct next station ", func() {
				So(shortestPathTrip.currentNode.name, ShouldEqual, "C")
			})
		})

	})
}

func TestCalcDistToConn(t *testing.T) {
	Convey("Given a new shortest path trip object", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("C")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)
		Convey("When ask to calculate dist to connections from its current node", func() {
			shortestPathTrip.CalcDistToConn()
			Convey("It calcualte it correctly", func() {
				s1, _ := n.GetNode("B")
				s2, _ := n.GetNode("D")
				s3, _ := n.GetNode("E")
				So(shortestPathTrip.distanceToStation[s1], ShouldEqual, 5)
				So(shortestPathTrip.distanceToStation[s2], ShouldEqual, 5)
				So(shortestPathTrip.distanceToStation[s3], ShouldEqual, 7)
			})
		})
	})

	Convey("Given a shortest path object thats visted one other node than the origin", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("C")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)
		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()
		Convey("When ask to calculate dist to connections from its current node", func() {
			shortestPathTrip.CalcDistToConn()

			Convey("It calcualte it correctly", func() {
				s1, _ := n.GetNode("B")
				s2, _ := n.GetNode("D")
				s3, _ := n.GetNode("E")
				s4, _ := n.GetNode("C")

				//Could return one of either these two stations.
				// i.e this is non deterministic but both these results are correct.
				if shortestPathTrip.currentNode.name == "B" {
					So(shortestPathTrip.distanceToStation[s1], ShouldEqual, 5)
					So(shortestPathTrip.distanceToStation[s2], ShouldEqual, 5)
					So(shortestPathTrip.distanceToStation[s3], ShouldEqual, 7)
					So(shortestPathTrip.distanceToStation[s4], ShouldEqual, 9)
				}
				if shortestPathTrip.currentNode.name == "D" {
					So(shortestPathTrip.distanceToStation[s1], ShouldEqual, 5)
					So(shortestPathTrip.distanceToStation[s2], ShouldEqual, 5)
					So(shortestPathTrip.distanceToStation[s3], ShouldEqual, 7)
					So(shortestPathTrip.distanceToStation[s4], ShouldEqual, 13)
				}
			})
		})
	})
}

func TestCompleted(t *testing.T) {

	Convey("Given a new shortest path object that still has stations to visit", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("C")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)
		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()

		Convey("When we ask if it is completed", func() {
			result := shortestPathTrip.Completed()

			Convey("It returns false", func() {
				So(result, ShouldBeFalse)
			})
		})

	})

	Convey("Given a new shortest path object that has no more stations to visit", t, func() {
		n, _ := getSimpleTestNetworkOfTrains()
		org, _ := n.GetNode("A")
		dest, _ := n.GetNode("C")
		t := trip{org, dest}
		shortestPathTrip := NewShortestPathTrip(n, &t)
		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()

		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()

		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()

		shortestPathTrip.CalcDistToConn()
		shortestPathTrip.VisitNextStation()

		Convey("When we ask if it is completed", func() {
			result := shortestPathTrip.Completed()

			Convey("It returns true", func() {
				So(result, ShouldBeTrue)

			})
		})
	})

	Convey("Given a network and origin and destination trip org and dest the same", t, func() {
		n, _ := getTestNetworkOfTrains()
		org, _ := n.GetNode("B")
		dest, _ := n.GetNode("B")
		t := trip{org, dest}

		Convey("When we ask if it is completed", func() {

			shortestPathTrip := NewShortestPathTrip(n, &t)
			shortestPathTrip.VisitNextStation()

			result := shortestPathTrip.Completed()

			Convey("It returns false", func() {
				So(result, ShouldBeFalse)

			})
		})

	})

}

// These setup the initial state we expect in the test.
func getVisitedStations(network *Network, t *trip) map[*station]bool {
	r := make(map[*station]bool)
	r[t.from] = true
	return r
}

func getStationsDistanceFromOrigin(n []*station, t *trip) map[*station]int {
	r := make(map[*station]int)
	for _, s := range n {
		r[s] = 99999
	}
	r[t.from] = 0
	return r
}
