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

func TestGetTripsBetweenArray(t *testing.T) {

	SkipConvey("Given a network of stations", t, func() {
		//get network
		Convey("When we ask for trips between two stations with stop constraint", func() {
			//ask for C-C
			Convey("It returns the correct result", func() {
				//C-D-C (2 stops). and C-E-B-C (3 stops)
			})
		})
	})

	SkipConvey("Given a query between two stations that are not connected", t, func() {
		//get network
		Convey("When we ask for trips between two stations with stop constraint", func() {

			Convey("It returns an error", func() {
				//?
			})
		})
	})

}
