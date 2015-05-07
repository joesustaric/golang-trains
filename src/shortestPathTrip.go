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

//GetNextStation sdkf
func (s ShortestPathTrip) GetNextStation() *station {
	// unvisitedConn := make(map[*station]int)
	//
	// for s, visited := range visitedSet {
	// 	if visited {
	// 		for c := range s.connections {
	// 			if !visitedSet[c] {
	// 				unvisitedConn[c] = shortestDistToNode[c]
	// 			}
	// 		}
	// 	}
	// }
	//
	// var result *station
	// lowestDist := 9999
	// for st, d := range unvisitedConn {
	// 	if d <= lowestDist {
	// 		lowestDist = d
	// 		result = st
	// 	}
	// }
	// return result
	unvisitedConn := make(map[*station]int)
	for n, visited := range s.visitedStations {
		if visited {
			for c := range n.connections {
				if !s.visitedStations[c] {
					unvisitedConn[c] = s.distanceToStation[c]
				}
			}
		}
	}

	var result *station
	lowestDist := 99999
	for st, d := range unvisitedConn {
		if d <= lowestDist {
			lowestDist = d
			result = st
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
