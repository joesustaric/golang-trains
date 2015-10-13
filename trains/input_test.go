package trains

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseInput(t *testing.T) {

	Convey("Given a valid input file loaction on disk", t, func() {
		testFile := "../fixtures/test_input.txt"

		Convey("When ParseInput is called", func() {
			result := ParseInput(testFile)

			Convey("Then it returns an array of valid train routes", func() {
				expected := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}

				So(result, ShouldResemble, expected)
			})
		})
	})
}

func TestValidateInput(t *testing.T) {

	Convey("Given a valid connection string", t, func() {
		connection := "AB5"

		Convey("When ValidateInput is called", func() {
			result := validateInput(connection)

			Convey("Then it returns true", func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey("Given an invalid connection string", t, func() {
		connection := "A45D"

		Convey("When ValidateInput is called", func() {
			result := validateInput(connection)

			Convey("Then it returns false", func() {
				So(result, ShouldBeFalse)
			})
		})
	})
}

func TestValidateInputDistanceQ(t *testing.T) {

	Convey("Given an valid distance calculation input", t, func() {
		input := "A-B-C"

		Convey("When ValidateInputDistanceQ is called", func() {
			result := ValidateInputDistanceQ(input)

			Convey("Then it returns array", func() {
				So(result, ShouldResemble, []string{"A", "B", "C"})
			})
		})
	})
}
