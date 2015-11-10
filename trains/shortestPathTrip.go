package trains

// ShortestPathTrip struct
type ShortestPathTrip struct {
	visitedSet   map[*Station]int
	unVisitedSet map[*Station]int
	Completed    bool
	originalTrip *trip
	currentNode  *Station
}

// INFINITY crap representation of infinity concept
var INFINITY = 99999

// NewShortestPathTrip new object
func NewShortestPathTrip(n *Network, t *trip) (*ShortestPathTrip, error) {
	result := &ShortestPathTrip{
		visitedSet:   initVisitedSet(t),
		unVisitedSet: initUnVisitedSet(n, t),
		Completed:    false,
		originalTrip: t,
		currentNode:  t.from}

	return result, nil
}

func initUnVisitedSet(n *Network, t *trip) map[*Station]int {
	result := make(map[*Station]int)
	for _, station := range n.stations {
		if stationShouldBeIncluded(station, t) {
			result[station] = INFINITY
		}
	}
	return result
}

func stationShouldBeIncluded(s *Station, t *trip) bool {
	notEqualToOrg := s != t.from
	orgAndDestEqual := s == t.from && s == t.to
	if orgAndDestEqual {
		return true
	}
	return notEqualToOrg
}

func initVisitedSet(trip *trip) map[*Station]int {
	result := make(map[*Station]int)
	result[trip.from] = 0
	return result
}

// CalcDistToConnectionsAndVisitNext blah
func (spt *ShortestPathTrip) CalcDistToConnectionsAndVisitNext() {
	if !spt.Completed {
		spt.calculateDistanceToCurrentStationConn()
		spt.setNextNodeToVisit()
		spt.addCurrentToVisitedAndRemoveFromUnVisited()
		spt.figureOutIfTripIsCompleted()
	}
}

func (spt *ShortestPathTrip) calculateDistanceToCurrentStationConn() {
	for c := range spt.currentNode.connections {
		newDist := spt.visitedSet[spt.currentNode] + spt.currentNode.GetDistanceTo(c)
		if newDist < spt.unVisitedSet[c] {
			spt.unVisitedSet[c] = newDist
		}
	}
}

func (spt *ShortestPathTrip) setNextNodeToVisit() {

	nextNode := &Station{}
	shortestDist := INFINITY

	for node := range spt.visitedSet {
		for conn := range node.connections {
			if _, ok := spt.unVisitedSet[conn]; ok {
				if node.GetDistanceTo(conn) < shortestDist {
					nextNode = conn
					shortestDist = node.GetDistanceTo(conn)
				}
			}
		}
	}
	spt.currentNode = nextNode
}

func (spt *ShortestPathTrip) addCurrentToVisitedAndRemoveFromUnVisited() {
	dist, _ := spt.unVisitedSet[spt.currentNode]
	spt.visitedSet[spt.currentNode] = dist
	delete(spt.unVisitedSet, spt.currentNode)
}

func (spt *ShortestPathTrip) figureOutIfTripIsCompleted() {
	if spt.currentNode == spt.originalTrip.to {
		spt.Completed = true
	}
}
