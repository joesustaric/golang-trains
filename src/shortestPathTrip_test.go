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

			Convey("Then it returns no error", func() {

				So(err, ShouldBeNil)

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
