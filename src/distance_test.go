package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistanceCalc(t *testing.T) {

	Convey("Given a Network of connecting stations", t, func() {
		network, _ := getTestNetworkOfTrains()

		Convey("When we ask for the distance from A-B-C", func() {
			query := []string{"A", "B", "C"}
			result := TotalDistance(network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 9)
			})
		})

		Convey("When we ask for the distance from A-D", func() {
			query := []string{"A", "D"}
			result := TotalDistance(network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When we ask for the distance from A-D-C", func() {
			query := []string{"A", "D", "C"}
			result := TotalDistance(network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 13)
			})
		})

		Convey("When we ask for the distance from A-E-B-C-D", func() {
			query := []string{"A", "E", "B", "C", "D"}
			result := TotalDistance(network, query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 22)
			})
		})

		Convey("When we ask for the distance to a station it does not connect to", func() {
			query := []string{"A", "E", "D"}
			result := TotalDistance(network, query)
			Convey("It returns 0", func() {
				So(result, ShouldEqual, 0)
			})
		})

	})
}
