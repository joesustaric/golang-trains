package trains

//ShortestPathTrip description
type ShortestPathTrip struct {
	network           *Network
	originalTrip      *trip
	visitedStations   map[*station]bool
	distanceToStation map[*station]int
}

//NewShortestPathTrip description
func NewShortestPathTrip(n *Network, t *trip) *ShortestPathTrip {
	visitedStations := createDefaultVisitedStationsMap(n, t)
	distToStation := initDistanceToStations(visitedStations)
	return &ShortestPathTrip{n, t, visitedStations, distToStation}
}

func createDefaultVisitedStationsMap(n *Network, t *trip) map[*station]bool {
	r := make(map[*station]bool)
	for _, s := range n.nodes {
		r[s] = false
	}
	r[t.from] = true
	return r
}

func initDistanceToStations(visitedSet map[*station]bool) map[*station]int {
	r := make(map[*station]int)
	for s, v := range visitedSet {
		if v {
			r[s] = 0
		} else {
			r[s] = 99999
		}

	}
	return r
}
