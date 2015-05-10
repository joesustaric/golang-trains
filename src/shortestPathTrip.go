package trains

import "fmt"

//this the representation of infinity
const INFINITY = 9999

//ShortestPathTrip description
type ShortestPathTrip struct {
	network           *Network
	originalTrip      *trip
	currentNode       *station
	visitedStations   map[*station]bool
	distanceToStation map[*station]int
}

//NewShortestPathTrip description
func NewShortestPathTrip(n *Network, t *trip) *ShortestPathTrip {
	visitedStations := createDefaultVisitedStationsMap(n, t)
	distToStation := initDistanceToStations(visitedStations, n)
	return &ShortestPathTrip{network: n, originalTrip: t, currentNode: t.from, visitedStations: visitedStations, distanceToStation: distToStation}
}

//GetNextStation todo
func (s ShortestPathTrip) GetNextStation() *station {

	unvisitedConn := s.createUnvisitedConnectionsMap(s.currentNode)

	var result *station
	lowestDist := INFINITY
	fmt.Println(len(unvisitedConn))
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

func (s ShortestPathTrip) createUnvisitedConnectionsMap(currentStation *station) map[*station]int {
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
		r[s] = 99999
	}

	for s := range visitedSet {
		r[s] = 0
	}
	return r
}
