package trains

import "fmt"

var MINIMUM_JOURNEY_NUMBER = 2

// TotalDistance returns the sum of the total distance between the given station
// journey.
func TotalDistance(n *Network, journey []string) (int, error) {
	if err := journeyValidation(n, journey); err != nil {
		return 0, err
	}

	return calculateTotalDistance(n, journey)
}

func calculateTotalDistance(n *Network, journey []string) (int, error) {
	totalDist, firstStation, theRestOfTheTrip := 0, journey[0], journey[1:]

	from := n.GetStation(firstStation)

	for _, destName := range theRestOfTheTrip {
		to := n.GetStation(destName)

		if canGetToNextFromCurrent(from, to) {
			totalDist += getDistance(from, to)
			from = to
		} else {
			return 0, fmt.Errorf("NO SUCH ROUTE")
		}
	}

	return totalDist, nil
}

func journeyValidation(n *Network, journey []string) error {
	if len(journey) < MINIMUM_JOURNEY_NUMBER {
		return fmt.Errorf("Journey must have at least 2 stations")
	}

	for _, st := range journey {
		if s := n.GetStation(st); s == nil {
			return fmt.Errorf("Station in the journey does not exist")
		}
	}
	return nil
}

func getDistance(o, d *Station) int {
	distance := 0
	for c, dist := range o.connections {
		if c.name == d.name {
			distance = dist
			break
		}
	}
	return distance
}

func canGetToNextFromCurrent(current, next *Station) bool {
	result := false
	for c := range current.connections {
		if c == next {
			result = true
			break
		}
	}
	return result
}
