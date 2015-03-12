package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlay(t *testing.T) {

	Convey("Given a station name 'F'", t, func() {
		name := "F"

		Convey("When we create a new station", func() {
			station := station{name}

			Convey("The station should be created correctly", func() {
				So(station.name, ShouldEqual, "F")
			})

		})
	})
}
