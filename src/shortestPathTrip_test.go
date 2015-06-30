package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewShortestPathTrip(t *testing.T) {

	SkipConvey("Given a valid test network of trains and a trip that is impossible", t, func() {

		Convey("When we ask for a NewShortestTrip Obejct", func() {

			Convey("Then it returns an error", func() {

			})
		})
	})

	Convey("Given a valid test network of trains and a valid trip to 2 different destinations", t, func() {

		network, stations := getSimplerTestNetworkOfTrains()
		orig, _ := network.GetNode("A")
		dest, _ := network.GetNode("B")
		testTrip := &trip{from: orig, to: dest}

		Convey("When we ask for a NewShortestTrip Obejct", func() {

			shortestPathTrip, err := NewShortestPathTrip(network, testTrip)

			Convey("Then it initalises the visited set correctly", func() {
				dist, ok := shortestPathTrip.visitedSet[orig]
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 0)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 1)
			})

			Convey("Then it initalises the unvisited set correctly", func() {
				for _, stn := range stations {
					if stn != testTrip.from || stn != testTrip.to {
						dist, ok := shortestPathTrip.unVisitedSet[stn]
						So(ok, ShouldBeTrue)
						So(dist, ShouldEqual, INFINITY)
					}
				}
			})

			Convey("Then it initalises the original trip properly", func() {
				So(shortestPathTrip.originalTrip, ShouldEqual, testTrip)
			})

			Convey("Then it marks the object as trip not completed", func() {
				So(shortestPathTrip.completed, ShouldBeFalse)
			})

			Convey("Then it assigned the from node as the current node", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, testTrip.from)
			})

			Convey("Then it returns no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a valid test network of trains and a valid trip to and from the same destination", t, func() {

		network, stations := getSimplerTestNetworkOfTrains()
		orig, _ := network.GetNode("B")
		dest := orig
		testTrip := &trip{from: orig, to: dest}

		Convey("When we ask for a NewShortestTrip Obejct", func() {

			shortestPathTrip, err := NewShortestPathTrip(network, testTrip)

			Convey("Then it initalises the visited set correctly", func() {
				dist, ok := shortestPathTrip.visitedSet[orig]
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 0)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 1)
			})

			Convey("Then it initalises the unvisited set correctly", func() {
				for _, stn := range stations {
					dist, ok := shortestPathTrip.unVisitedSet[stn]
					So(ok, ShouldBeTrue)
					So(dist, ShouldEqual, INFINITY)
				}
			})

			Convey("Then it initalises the original trip properly", func() {
				So(shortestPathTrip.originalTrip, ShouldEqual, testTrip)
			})

			Convey("Then it marks the object as trip not completed", func() {
				So(shortestPathTrip.completed, ShouldBeFalse)
			})

			Convey("Then it assigned the from node as the current node", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, testTrip.from)
			})

			Convey("Then it returns no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an initalised shortest path trip with valid trip to different to and from destinations one hop apart", t, func() {

		network, _ := getSimplerTestNetworkOfTrains()
		orig, _ := network.GetNode("A")
		dest, _ := network.GetNode("B")
		testTrip := &trip{from: orig, to: dest}
		shortestPathTrip, _ := NewShortestPathTrip(network, testTrip)

		Convey("When we ask to calculate the distance to the current nodes connections and visit next", func() {

			shortestPathTrip.CalcDistToConnectionsAndVisitNext()

			Convey("Then it calculates the correct distance to the unvisited nodes", func() {
				b, _ := network.GetNode("B")
				c, _ := network.GetNode("C")
				d, _ := network.GetNode("D")
				So(shortestPathTrip.unVisitedSet[b], ShouldEqual, 1)
				So(shortestPathTrip.unVisitedSet[c], ShouldEqual, 4)
				So(shortestPathTrip.unVisitedSet[d], ShouldEqual, INFINITY)
			})

			Convey("Then it makes the next unvisited node with the shortest distance the current", func() {
				b, _ := network.GetNode("B")
				So(shortestPathTrip.currentNode, ShouldEqual, b)
			})

			Convey("Then it marks the calculation as completed", func() {
				So(shortestPathTrip.completed, ShouldBeTrue)
			})

		})
	})

}

// Convey("Given", t, func() {
//
// 	Convey("When", func() {
//
// 		Convey("Then", func() {
//
// 		})
// 	})
// })
