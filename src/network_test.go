package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddNodeAndGetNode(t *testing.T) {

	Convey("Given a station name 'F'", t, func() {
		name := "F"
		s := &station{name, map[*station]int{}}
		Convey("When we add it to the network", func() {
			network := network{map[string]*station{}}
			network.AddNode(s)
			Convey("We can get it back", func() {
				result := network.GetNode(name)
				So(result, ShouldEqual, s)
			})

		})
	})
}
