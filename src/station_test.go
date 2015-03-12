package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStation(t *testing.T) {

	Convey("Given a station name 'F'", t, func() {
		name := "F"

		Convey("When we create a new station", func() {
			station := station{name, map[*station]int{}}

			Convey("The station should be created correctly", func() {
				So(station.name, ShouldEqual, "F")
			})

		})
	})

	Convey("Given a station object", t, func() {
		name := "Foo"
		foo := station{name, map[*station]int{}}

		Convey("When we add a station it connects to", func() {
			conectionName := "Bar"
			bar := station{conectionName, map[*station]int{}}
			distance := 5
			foo.AddConnection(&bar, distance)

			Convey("The station keeps a refrence to the connection with its distance", func() {
				So(foo.GetConnection(conectionName), ShouldResemble, bar)
				So(foo.GetDistanceTo(&bar), ShouldEqual, 5)
			})
		})
	})
}
