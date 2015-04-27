package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetShortestRoute(t *testing.T) {

	Convey("Given a network of stations", t, func() {
		network := getTestNetworkOfTrains()

		Convey("When we ask for the shortest path from A to C", func() {
			result := GetShortestRouteDistance(network, "A", "C")
			Convey("It returns the shortest path", func() {
				So(result, ShouldEqual, 9)
			})
		})

		// Convey("When we ask for the shortest path from B to B", func() {
		// 	result := GetShortestRouteDistance(network, "B", "B")
		// 	Convey("It returns the shortest path", func() {
		// 		So(result, ShouldEqual, 9)
		// 	})
		// })

	})
}
