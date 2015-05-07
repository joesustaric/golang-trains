package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTripsBetween(t *testing.T) {

	Convey("Given a network of stations", t, func() {
		network, _ := getTestNetworkOfTrains()

		Convey("When we ask for trips between C-C with stop constraint of 3", func() {
			t := tripHopsQuery{origin: "C", destination: "C", maxHops: 3}
			result := GetNumberOfRoutes(network, t)
			Convey("It returns an array with all the correct routes", func() {
				So(result, ShouldEqual, 2)
			})
		})

		Convey("When we ask for trips between A-C with stop constraint of 4", func() {
			t := tripHopsQuery{origin: "A", destination: "C", maxHops: 4}
			result := GetNumberOfRoutes(network, t)
			Convey("It returns an array with all the correct routes", func() {
				So(result, ShouldEqual, 3)
			})
		})
	})

}
