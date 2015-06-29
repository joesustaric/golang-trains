package trains

// ShortestPathTrip struct
type ShortestPathTrip struct {
	visitedSet   map[*station]int
	unVisitedSet map[*station]int
	completed    bool
	originalTrip *trip
}

// INFINITY crap representation of infinity concept
var INFINITY = 99999

// NewShortestPathTrip new object
func NewShortestPathTrip(n *Network, t *trip) (*ShortestPathTrip, error) {
	result := &ShortestPathTrip{
		visitedSet:   initVisitedSet(t),
		unVisitedSet: initUnVisitedSet(n, t),
		completed:    false,
		originalTrip: t}

	return result, nil
}

func initUnVisitedSet(n *Network, t *trip) map[*station]int {
	result := make(map[*station]int)
	for _, stn := range n.nodes {
		if stn != t.from || stn != t.to {
			result[stn] = INFINITY
		}
	}
	return result
}

func initVisitedSet(trip *trip) map[*station]int {
	result := make(map[*station]int)
	result[trip.from] = 0
	return result
}
