package trains

import "fmt"

// visited set
//
// start at origin
// calculated distance to all connecting nodes from current
// (if the calc is < current distance (initially infinity)) assign it to its value
// go to the lowest connecting dist node
// add it to the visited set
// get all unvisited connections and recalc distance (if lower etc)
// do this for the next lowest connection

//GetShortestRoute will get the shortest path between two stations
func GetShortestRoute(n Network, origin, destination string) int {

	shortestDistToNode := make(map[*station]int)

	//add all the paths and set their dist to 999 (infinity)
	for _, s := range n.nodes {
		shortestDistToNode[s] = 999
	}

	visitedSet := make(map[*station]bool)
	//add all nodes to the visited set and flag them as not visited
	for _, s := range n.nodes {
		visitedSet[s] = false
	}

	//get origin node
	org, ok := n.GetNode(origin)
	if !ok {
		return 0
	}
	//get destination node
	des, ok := n.GetNode(destination)
	if !ok {
		return 0
	}

	//mark origin as visited and distance of 0
	visitedSet[org] = true
	shortestDistToNode[org] = 0
	//calc distance to all connecting nodes from current node
	calcDistToConnections(org, shortestDistToNode, visitedSet)
	//get next node to visit
	nextStation := getNextStationToVisit(org, shortestDistToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if visitedSet[des] {
		return shortestDistToNode[des]
	}

	//calc distance to all connecting nodes from current node
	calcDistToConnections(nextStation, shortestDistToNode, visitedSet)
	//get next node to visit
	nextStation = getNextStationToVisit(nextStation, shortestDistToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if visitedSet[des] {
		return shortestDistToNode[des]
	}

	//calc distance to all connecting nodes from current node
	calcDistToConnections(nextStation, shortestDistToNode, visitedSet)
	//get next node to visit
	nextStation = getNextStationToVisit(nextStation, shortestDistToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if visitedSet[des] {
		return shortestDistToNode[des]
	}

	//calc distance to all connecting nodes from current node
	calcDistToConnections(nextStation, shortestDistToNode, visitedSet)
	//get next node to visit
	nextStation = getNextStationToVisit(nextStation, shortestDistToNode, visitedSet)
	//add next station to visited set
	visitedSet[nextStation] = true
	//have we reached the destination?
	if visitedSet[des] {
		return shortestDistToNode[des]
	}

	return 1234
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
		fmt.Println((shortestDistToNode[currentNode] + distToConn))

		if !visitedSet[c] {
			if (shortestDistToNode[currentNode] + distToConn) <= shortestDistToNode[c] {
				shortestDistToNode[c] = (shortestDistToNode[currentNode] + distToConn)
			}
		}
	}
}
