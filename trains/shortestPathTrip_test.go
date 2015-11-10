package trains

import (
	"fmt"
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
		orig, dest := network.GetStation("A"), network.GetStation("B")
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
					if stn != testTrip.from {
						dist, ok := shortestPathTrip.unVisitedSet[stn]
						So(ok, ShouldBeTrue)
						So(dist, ShouldEqual, INFINITY)
					}
				}
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 3)
			})

			Convey("Then it initalises the original trip properly", func() {
				So(shortestPathTrip.originalTrip, ShouldEqual, testTrip)
			})

			Convey("Then it marks the object as trip not completed", func() {
				So(shortestPathTrip.Completed, ShouldBeFalse)
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
		orig := network.GetStation("B")
		testTrip := &trip{from: orig, to: orig}

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
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 4)
			})

			Convey("Then it initalises the original trip properly", func() {
				So(shortestPathTrip.originalTrip, ShouldEqual, testTrip)
			})

			Convey("Then it marks the object as trip not completed", func() {
				So(shortestPathTrip.Completed, ShouldBeFalse)
			})

			Convey("Then it assigned the from node as the current node", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, testTrip.from)
			})

			Convey("Then it returns no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a valid trip to different to and from destinations one hop apart", t, func() {

		network, _ := getSimplerTestNetworkOfTrains()
		orig, dest := network.GetStation("A"), network.GetStation("B")
		testTrip := &trip{from: orig, to: dest}
		shortestPathTrip, _ := NewShortestPathTrip(network, testTrip)

		Convey("When we ask to calculate the distance to the current nodes connections and visit next", func() {

			shortestPathTrip.CalcDistToConnectionsAndVisitNext()

			Convey("Then removes the current node form the unVisited set", func() {
				_, ok := shortestPathTrip.unVisitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeFalse)
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 2)
			})

			Convey("Then it calculates the correct distance to the unvisited nodes", func() {
				c, d := network.GetStation("C"), network.GetStation("D")
				So(shortestPathTrip.unVisitedSet[c], ShouldEqual, 4)
				So(shortestPathTrip.unVisitedSet[d], ShouldEqual, INFINITY)
			})

			Convey("Then it makes the next unvisited node with the shortest distance the current", func() {
				b := network.GetStation("B")
				So(shortestPathTrip.currentNode, ShouldEqual, b)
			})

			Convey("Then it adds the current node to the visited set with the correct distance", func() {
				dist, ok := shortestPathTrip.visitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 1)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 2)
			})

			Convey("Then it marks the calculation as completed", func() {
				So(shortestPathTrip.Completed, ShouldBeTrue)
			})

		})
	})

	Convey("Given a valid trip to the same to and from destination", t, func() {

		network, _ := getSimplerTestNetworkOfTrains()
		orig, dest := network.GetStation("A"), network.GetStation("A")
		testTrip := &trip{from: orig, to: dest}
		shortestPathTrip, _ := NewShortestPathTrip(network, testTrip)

		Convey("When we ask to calculate the distance to the current nodes connections and visit next once", func() {

			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			a, b := network.GetStation("A"), network.GetStation("B")
			c, d := network.GetStation("C"), network.GetStation("D")

			Convey("Then removes the current node form the unVisited set", func() {
				_, ok := shortestPathTrip.unVisitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeFalse)
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 3)
			})

			Convey("Then it calculates the correct distance to the unvisited nodes", func() {
				So(shortestPathTrip.unVisitedSet[a], ShouldEqual, INFINITY)
				So(shortestPathTrip.unVisitedSet[c], ShouldEqual, 4)
				So(shortestPathTrip.unVisitedSet[d], ShouldEqual, INFINITY)
			})

			Convey("Then it adds the current node to the visited set with the correct distance", func() {
				dist, ok := shortestPathTrip.visitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 1)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 2)
			})

			Convey("Then it makes the next unvisited node with the shortest distance the current", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, b)
			})

			Convey("Then it marks the calculation as completed", func() {
				So(shortestPathTrip.Completed, ShouldBeFalse)
			})

		})

		Convey("When we ask to calculate the distance to the current nodes connections and visit next twice", func() {

			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			a, c, d := network.GetStation("A"), network.GetStation("C"), network.GetStation("D")

			Convey("Then removes the current node form the unVisited set", func() {
				_, ok := shortestPathTrip.unVisitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeFalse)
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 2)
			})

			Convey("Then it calculates the correct distance to the unvisited nodes", func() {
				So(shortestPathTrip.unVisitedSet[d], ShouldEqual, 6)
				So(shortestPathTrip.unVisitedSet[a], ShouldEqual, INFINITY)
			})

			Convey("Then it adds the current node to the visited set with the correct distance", func() {
				dist, ok := shortestPathTrip.visitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 3)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 3)
			})

			Convey("Then it makes the next unvisited node with the shortest distance the current", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, c)
			})

			Convey("Then it marks the calculation as completed", func() {
				So(shortestPathTrip.Completed, ShouldBeFalse)
			})

		})

		Convey("When we ask to calculate the distance to the current nodes connections and visit next thrice", func() {

			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			shortestPathTrip.CalcDistToConnectionsAndVisitNext()
			a, d := network.GetStation("A"), network.GetStation("D")

			Convey("Then removes the current node form the unVisited set", func() {
				_, ok := shortestPathTrip.unVisitedSet[shortestPathTrip.currentNode]
				So(ok, ShouldBeFalse)
				So(len(shortestPathTrip.unVisitedSet), ShouldEqual, 1)
			})

			Convey("Then it calculates the correct distance to the unvisited nodes", func() {
				So(shortestPathTrip.unVisitedSet[d], ShouldEqual, 6)
			})

			Convey("Then it adds the current node to the visited set with the correct distance", func() {
				dist, ok := shortestPathTrip.visitedSet[shortestPathTrip.currentNode]
				fmt.Println(shortestPathTrip.currentNode)
				So(ok, ShouldBeTrue)
				So(dist, ShouldEqual, 6)
				So(len(shortestPathTrip.visitedSet), ShouldEqual, 3)
			})

			Convey("Then it makes the next unvisited node with the shortest distance the current", func() {
				So(shortestPathTrip.currentNode, ShouldEqual, a)
			})

			Convey("Then it marks the calculation as completed", func() {
				So(shortestPathTrip.Completed, ShouldBeTrue)
			})
		})

	})
}
