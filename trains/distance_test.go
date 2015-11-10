package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistanceCalc(t *testing.T) {

	Convey("Given a Network of connecting stations", t, func() {
		network, _ := getTestNetworkOfTrains()

		//Question #1
		Convey("When we ask for the distance from A-B-C", func() {
			result, _ := TotalDistance(network, []string{"A", "B", "C"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 9)
			})
		})

		//Question #2
		Convey("When we ask for the distance from A-D", func() {
			result, _ := TotalDistance(network, []string{"A", "D"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 5)
			})
		})

		//Question #3
		Convey("When we ask for the distance from A-D-C", func() {
			result, _ := TotalDistance(network, []string{"A", "D", "C"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 13)
			})
		})

		//Question #4
		Convey("When we ask for the distance from A-E-B-C-D", func() {
			result, err := TotalDistance(network, []string{"A", "E", "B", "C", "D"})

			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 22)
			})
			Convey("It returns no error", func() {
				So(err, ShouldBeNil)
			})
		})

		//Question #5
		Convey("When we ask for the distance to a station it does not connect to", func() {
			_, err := TotalDistance(network, []string{"A", "E", "D"})

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "NO SUCH ROUTE")
			})
		})

		Convey("When we ask for a journey with only 1 station", func() {
			_, err := TotalDistance(network, []string{"A"})

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "at least 2 stations")
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
