package trains

import "errors"

//GetShortestRouteDistance will get the shortest path between two stations
func GetShortestRouteDistance(n *Network, origin, destination string) int {

	trip, err := createTrip(origin, destination, n)
	if err != nil {
		return 0
	}
	shortestDistToNode := createShortestVisitedSet(n, trip.from)
	visitedStation := createVisitedStationSet(n, trip.from)

	return getShortestDistance(trip, shortestDistToNode, visitedStation)
}

func createTrip(origin, destination string, n *Network) (*trip, error) {
	result := trip{}
	org, ok := n.GetNode(origin)
	if !ok {
		return nil, errors.New("origin does not exists")
	}
	des, ok := n.GetNode(destination)
	if !ok {
		return nil, errors.New("destination does not exists")
	}
	result.to = des
	result.from = org
	return &result, nil
}

func createVisitedStationSet(n *Network, origin *station) map[*station]bool {
	result := make(map[*station]bool)
	result[origin] = true //mark origin visited
	return result
}

func createShortestVisitedSet(n *Network, origin *station) map[*station]int {
	result := make(map[*station]int)
	//9999 = infinity representation
	for _, s := range n.nodes {
		result[s] = 99999
	}
	result[origin] = 0 //0 distance to get form origin to origin
	return result
}

func getShortestDistance(t *trip, shortestDistanceToNode map[*station]int, visitedSet map[*station]bool) int {
	//calc distance to all connecting nodes from current node
	calcDistToConnections(t.from, shortestDistanceToNode, visitedSet)
	//get next node to visit
	nextStation := getNextStationToVisit(t.from, shortestDistanceToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if !visitedSet[t.to] {
		getShortestDistance(&trip{nextStation, t.to}, shortestDistanceToNode, visitedSet)
	}

	return shortestDistanceToNode[t.to]
}

func getNextStationToVisit(current *station, shortestDistToNode map[*station]int, visitedSet map[*station]bool) *station {

	unvisitedConn := make(map[*station]int)

	for s, visited := range visitedSet {
		if visited {
			for c := range s.connections {
				if !visitedSet[c] {
					unvisitedConn[c] = shortestDistToNode[c]
				}
			}
		}
	}

	var result *station
	lowestDist := 9999
	for st, d := range unvisitedConn {
		if d <= lowestDist {
			lowestDist = d
			result = st
		}
	}
	return result
}

func calcDistToConnections(currentNode *station, shortestDistToNode map[*station]int, visitedSet map[*station]bool) {

	for c, distToConn := range currentNode.connections {

		if !visitedSet[c] {
			if (shortestDistToNode[currentNode] + distToConn) <= shortestDistToNode[c] {
				shortestDistToNode[c] = (shortestDistToNode[currentNode] + distToConn)
			}
		}
	}
}

func originAndDestinationInValid(org, dest *station) bool {
	if org == nil || dest == nil {
		return true
	}
	return false
}
