package trains

// TotalDistance returns the sum of the total distance between the given station
// journey.
func TotalDistance(n *Network, journey []string) int {
	totalDist := 0
	from, ok := n.GetNode(journey[0])
	if !ok {
		return 0
	}

	for _, destName := range journey[1:] {
		to, ok := n.GetNode(destName)
		if !ok {
			return 0
		}
		if canGetToNextFromCurrent(from, to) {
			totalDist += getDistance(from, to)
			from = to
		} else {
			return 0
		}
	}
	return totalDist
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
