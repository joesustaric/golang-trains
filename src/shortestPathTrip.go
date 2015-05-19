package trains

// This my representation of infinity for lack of a better way.
const INFINITY = 99999

// ShortestPathTrip this is the data structure which contains all the data
// needed to calculate the shortest path trip.
type ShortestPathTrip struct {
	network           *Network
	originalTrip      *trip
	currentNode       *station
	visitedStations   map[*station]bool
	distanceToStation map[*station]int
}

// NewShortestPathTrip this will initialise a new ShortestPathTrip object with
// all the correct default values
func NewShortestPathTrip(n *Network, t *trip) *ShortestPathTrip {
	visitedStations := createDefaultVisitedStationsMap(n, t)
	distToStation := initDistanceToStations(visitedStations, n)
	return &ShortestPathTrip{network: n,
		originalTrip: t, currentNode: t.from, visitedStations: visitedStations, distanceToStation: distToStation}
}

// GetNextStation returns the next station to visit based on Dijkstra's algorithm.
// It will return the next unvisited station to visit with the next lowest
// connection distance, based on the current node were at.
func (s *ShortestPathTrip) GetNextStation() *station {
	var result *station
	lowestDist := INFINITY
	for st, d := range s.currentNode.connections {
		if !s.visitedStations[st] {
			if d <= lowestDist {
				lowestDist = d
				result = st
			}
		}
	}
	return result
}

// CalcDistToConn will calculate the distance to all the connecting nodes form
// the current connecting node
func (s *ShortestPathTrip) CalcDistToConn() {
	for st, d := range s.currentNode.connections {
		if s.distanceToStation[st] == INFINITY {
			s.distanceToStation[st] = d
		} else {
			s.distanceToStation[st] = s.distanceToStation[st] + d
		}
	}
}

func (s *ShortestPathTrip) createUnvisitedConnectionsMap(currentStation *station) map[*station]int {
	unvisitedConn := make(map[*station]int)
	for st, dist := range currentStation.connections {
		for c := range st.connections {
			if !s.visitedStations[c] {
				unvisitedConn[c] = dist
			}
		}
	}
	return unvisitedConn
}

func createDefaultVisitedStationsMap(n *Network, t *trip) map[*station]bool {
	r := make(map[*station]bool)
	r[t.from] = true
	return r
}

func initDistanceToStations(visitedSet map[*station]bool, n *Network) map[*station]int {
	r := make(map[*station]int)
	for name := range n.nodes {
		s := n.nodes[name]
		r[s] = INFINITY
	}

	for s := range visitedSet {
		r[s] = 0
	}
	return r
}
