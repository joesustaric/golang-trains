package trains

import "fmt"

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

//GetNextStation sdkf
func (s ShortestPathTrip) GetNextStation() *station {

	unvisitedConn := make(map[*station]int)
	for st, dist := range s.currentNode.connections {
		for c := range st.connections {
			if !s.visitedStations[c] {
				unvisitedConn[c] = dist
			}
		}
	}

	var result *station
	lowestDist := 99999
	fmt.Println(len(unvisitedConn))
	for st, d := range s.currentNode.connections {
		if !s.visitedStations[st] {
			if d <= lowestDist {
				fmt.Println("--", lowestDist, "+", st.name, "d=", d)
				lowestDist = d
				result = st
				fmt.Println("!!", lowestDist, "+", st.name, "d=", d, "result=", result.name)
			}
		}

	}
	return result
}

func createDefaultVisitedStationsMap(n *Network, t *trip) map[*station]bool {
	r := make(map[*station]bool)
	// for _, s := range n.nodes {
	// 	r[s] = false
	// }
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
