package io

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInput(t *testing.T) {
	Convey("Given a valid input file loaction on disk", t, func() {
		testFile := "../../fixtures/test_input.txt"

		Convey("When ParseInput is called", func() {
			result, _ := ParseInput(testFile)

			Convey("Then it returns an array of valid train routes", func() {
				expected := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
				So(result, ShouldResemble, expected)
			})
		})
	})

}
