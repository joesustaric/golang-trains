package trains

import (
	"errors"
	"fmt"
)

// TotalDistance returns the sum of the total distance between the given station
// journey.
func TotalDistance(n *Network, journey []string) (int, error) {
	if ok := journeyStationsExist(n, journey); !ok {
		return 0, errors.New("Station in the journey does not exist")
	}
	totalDist := 0

	if from, ok := n.GetNode(journey[0]); ok {
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
	return 0, fmt.Errorf("Error something")
}

func journeyStationsExist(n *Network, journey []string) bool {
	for _, st := range journey {
		if _, ok := n.GetNode(st); !ok {
			return false
		}
	}
	return true
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
