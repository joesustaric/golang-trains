package trains

// ShortestPathTrip struct
type ShortestPathTrip struct {
	visitedSet   map[*station]int
	unVisitedSet map[*station]int
	Completed    bool
	originalTrip *trip
	currentNode  *station
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

func initUnVisitedSet(n *Network, t *trip) map[*station]int {
	result := make(map[*station]int)
	for _, stn := range n.nodes {
		if stationShouldBeIncluded(stn, t) {
			result[stn] = INFINITY
		}
	}
	return result
}

func stationShouldBeIncluded(s *station, t *trip) bool {
	notEqualToOrg := s != t.from
	orgAndDestEqual := s == t.from && s == t.to
	if orgAndDestEqual {
		return true
	}
	return notEqualToOrg
}

func initVisitedSet(trip *trip) map[*station]int {
	result := make(map[*station]int)
	result[trip.from] = 0
	return result
}

// CalcDistToConnectionsAndVisitNext blah
func (spt *ShortestPathTrip) CalcDistToConnectionsAndVisitNext() {
	// if not completed
	// go through all connections for current node
	// add node shortest distance calc to the connection
	// if lower then current value for that node
	// make it the new value

	//go through unvisited set find the node with shortest dist
	//remove it from unvisited
	//add it to visited
	//make it current node

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

	nextNode := &station{}
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
