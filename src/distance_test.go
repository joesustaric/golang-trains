package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistanceCalc(t *testing.T) {

	Convey("Given 'A' station with 1 connection", t, func() {
		a := station{"A", map[*station]int{}}
		b := station{"B", map[*station]int{}}
		a.AddConnection(&b, 5)
		query := []string{"A", "B"}
		Convey("When we ask for the distance from A-B", func() {
			result := TotalDistance(query)
			Convey("It returns the correct distance", func() {
				So(result, ShouldEqual, 5)
			})

		})
	})

	// Convey("Given 'A' station with 2 connections", t, func() {
	// 	a := station{"A", map[*station]int{}}
	// 	b := station{"B", map[*station]int{}}
	// 	c := station{"C", map[*station]int{}}
	// 	a.AddConnection(&b, 5)
	// 	b.AddConnection(&c, 4)
	//
	// 	Convey("When we ask for the distance from A-B-C", func() {
	// 		result := TotalDistance(a, b, c)
	// 		Convey("It returns the correct distance", func() {
	// 			So(result, ShouldEqual, 9)
	// 		})
	//
	// 	})
	// })

}
