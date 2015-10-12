package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStation(t *testing.T) {

	Convey("Given a station name 'F'", t, func() {
		name := "F"
		Convey("When we create a new station", func() {
			s := NewStation(name)
			Convey("The station should be created correctly", func() {
				So(s.name, ShouldEqual, "F")
			})
		})
	})
}

func TestAddConnection(t *testing.T) {

	Convey("Given a station object", t, func() {
		name := "Foo"
		foo := NewStation(name)
		Convey("When we add a station it connects to", func() {
			connectionName := "Bar"
			bar := NewStation(connectionName)
			distance := 5
			foo.AddConnection(bar, distance)
			Convey("The station keeps a refrence to the connection with its distance", func() {
				So(foo.GetConnection(connectionName), ShouldResemble, bar)
				So(foo.GetDistanceTo(bar), ShouldEqual, 5)
			})
		})
	})
}
