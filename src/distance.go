package trains

import "errors"

// TotalDistance returns the sum of the total distance between the given station
// journey.
func TotalDistance(n *Network, journey []string) (int, error) {
	totalDist := 0
	ok := journeyStationsExist(n, journey)
	if !ok {
		return 0, errors.New("Station in the journey does not exist")
	}

	from, _ := n.GetNode(journey[0])
	for _, destName := range journey[1:] {
		to, _ := n.GetNode(destName)

		if canGetToNextFromCurrent(from, to) {
			totalDist += getDistance(from, to)
			from = to
		} else {
			return 0, errors.New("Can't get to a connection " + to.name)
		}
	}
	return totalDist, nil
}

func journeyStationsExist(n *Network, journey []string) bool {
	for _, st := range journey {
		_, ok := n.GetNode(st)
		if !ok {
			return false
		}
	}
	return true
}

func getDistance(o, d *station) int {
	for c, dist := range o.connections {
		if c.name == d.name {
			return dist
		}
	}
	return 0
}

func canGetToNextFromCurrent(current *station, next *station) bool {
	result := false
	for c := range current.connections {
		if c == next {
			result = true
		}
	}
	return result
}
