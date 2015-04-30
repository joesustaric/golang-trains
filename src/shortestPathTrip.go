package trains

type shortestPathTrip struct {
	network         *Network
	originalTrip    *trip
	visitedStations map[*station]bool
}

func NewShortestPathTrip(n *Network, t *trip) *shortestPathTrip {
	visitedStations := createDefaultVisitedStationsMap(n)
	return &shortestPathTrip{n, t, visitedStations}
}

func createDefaultVisitedStationsMap(n *Network) map[*station]bool {
	r := make(map[*station]bool)

	for _, s := range n.nodes {
		r[s] = false
	}

	return r
}
