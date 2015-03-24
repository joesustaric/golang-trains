package trains

// TotalDistance description
func TotalDistance(n Network, input []string) int {
	totalDist := 0
	from, ok := n.GetNode(input[0])
	if !ok {
		return 0
	}

	for _, destName := range input[1:] {
		to, ok := n.GetNode(destName)
		if !ok {
			return 0
		}
		totalDist += getDistance(from, to)
		from = to
	}

	return totalDist
}

func getDistance(o *station, d *station) int {
	for c, dist := range o.connections {
		if c.name == d.name {
			return dist
		}
	}
	return 0
}
