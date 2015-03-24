package trains

// TotalDistance description
func TotalDistance(n network, input []string) int {
	totalDist := 0
	origin := n.GetNode(input[0])
	for c, dist := range origin.connections {
		if c.name == input[1] {
			totalDist += dist
		}
	}
	return totalDist
}
