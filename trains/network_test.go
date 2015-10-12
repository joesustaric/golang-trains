package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddNodeAndGetNode(t *testing.T) {

	Convey("Given a station name 'F'", t, func() {
		name := "F"
		s := NewStation(name)
		Convey("When we add it to the Network", func() {
			Network := NewNetwork()
			Network.AddNode(s)
			Convey("We can get it back", func() {
				result, _ := Network.GetNode(name)
				So(result, ShouldEqual, s)
			})
		})
	})

	Convey("Given an empty Network", t, func() {
		Network := NewNetwork()
		Convey("When we ask for a station", func() {
			station, ok := Network.GetNode("A")
			Convey("It returns nil", func() {
				So(station, ShouldBeNil)
			})
			Convey("It returns not found", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
