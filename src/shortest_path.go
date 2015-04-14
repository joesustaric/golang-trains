package trains

//GetShortestRouteDistance will get the shortest path between two stations
func GetShortestRouteDistance(n Network, origin, destination string) int {

	org, dest := getOriginAndDestinationNodes(origin, destination, n)
	if org == nil || dest == nil {
		return 0
	}
	shortestDistToNode := createShortestVisitedSet(n, org)
	visitedStation := createVisitedStationSet(n, org)

	return getShortestDistance(org, dest, shortestDistToNode, visitedStation)
}

func getOriginAndDestinationNodes(origin, destination string, n Network) (*station, *station) {

	org, ok := n.GetNode(origin)
	if !ok {
		return nil, nil
	}
	des, ok := n.GetNode(destination)
	if !ok {
		return nil, nil
	}
	return org, des
}

func createVisitedStationSet(n Network, origin *station) map[*station]bool {
	result := make(map[*station]bool)
	for _, s := range n.nodes {
		result[s] = false
	}
	result[origin] = true //mark origin visited
	return result
}

func createShortestVisitedSet(n Network, origin *station) map[*station]int {
	result := make(map[*station]int)
	//999 = infinity representation
	for _, s := range n.nodes {
		result[s] = 9999
	}
	result[origin] = 0 //0 distance to get form origin to origin
	return result
}

func getShortestDistance(origin, destination *station, shortestDistanceToNode map[*station]int, visitedSet map[*station]bool) int {
	//calc distance to all connecting nodes from current node
	calcDistToConnections(origin, shortestDistanceToNode, visitedSet)
	//get next node to visit
	nextStation := getNextStationToVisit(origin, shortestDistanceToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if !visitedSet[destination] {
		getShortestDistance(nextStation, destination, shortestDistanceToNode, visitedSet)
	}
	return shortestDistanceToNode[destination]
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
