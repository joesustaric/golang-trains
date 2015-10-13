package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistanceCalc(t *testing.T) {

	Convey("Given a Network of connecting stations", t, func() {
		network, _ := getTestNetworkOfTrains()

		Convey("When we ask for the distance from A-B-C", func() {
			result, _ := TotalDistance(network, []string{"A", "B", "C"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 9)
			})
		})

		Convey("When we ask for the distance from A-D", func() {
			result, _ := TotalDistance(network, []string{"A", "D"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When we ask for the distance from A-D-C", func() {
			result, _ := TotalDistance(network, []string{"A", "D", "C"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 13)
			})
		})

		Convey("When we ask for the distance from A-E-B-C-D", func() {
			result, _ := TotalDistance(network, []string{"A", "E", "B", "C", "D"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 22)
			})
		})

		Convey("When we ask for the distance to a station it does not connect to", func() {
			_, err := TotalDistance(network, []string{"A", "E", "D"})

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When we ask for the distance to a station it does not exist", func() {
			_, err := TotalDistance(network, []string{"A", "E", "B", "C", "Z"})

			Convey("It returns the correct distance", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

}
