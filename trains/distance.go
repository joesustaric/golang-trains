package trains

import "fmt"

// TotalDistance returns the sum of the total distance between the given station
// journey.
func TotalDistance(n *Network, journey []string) (int, error) {
	if err := journeyStationsExist(n, journey); err != nil {
		return 0, err
	}

	totalDist := 0

	//needs refactoring here
	if from, ok := n.GetNode(journey[0]); ok {
		for _, destName := range journey[1:] {

			if to, ok := n.GetNode(destName); ok {

				if canGetToNextFromCurrent(from, to) {
					totalDist += getDistance(from, to)
					from = to
				} else {
					return 0, fmt.Errorf("Can't get to a connection %s", to.name)
				}
			}
		}
	}
	return totalDist, nil
}

func journeyStationsExist(n *Network, journey []string) error {
	for _, st := range journey {
		if _, ok := n.GetNode(st); !ok {
			return fmt.Errorf("Station in the journey does not exist")
		}
	}
	return nil
}

func getDistance(o, d *Station) int {
	for c, dist := range o.connections {
		if c.name == d.name {
			return dist
		}
	}
	return 0
}

func canGetToNextFromCurrent(current, next *Station) bool {
	result := false
	for c := range current.connections {
		if c == next {
			result = true
		}
	}
	return result
}
