package trains

type tripHopsQuery struct {
	origin      string
	destination string
	maxHops     int
}

type trip struct {
	from *station
	to   *station
}

const firstHop int = 1
const singleHop int = 1

// GetNumberOfRoutes returns the number of routes in a network with a hops query.
func GetNumberOfRoutes(n *Network, thq tripHopsQuery) int {
	origin, ok := n.GetNode(thq.origin)
	if !ok {
		return 0
	}
	dest, ok := n.GetNode(thq.destination)
	if !ok {
		return 0
	}
	tripNodes := trip{from: origin, to: dest}
	return numberOfRoutes(tripNodes, firstHop, thq.maxHops)
}

func numberOfRoutes(t trip, hopsTravelled, maxHops int) int {
	result := 0
	for conn := range t.from.connections {
		if conn == t.to || result < maxHops {
			result++
		} else if hopsTravelled < maxHops {
			connTripNodes := trip{from: conn, to: t.to}
			result += numberOfRoutes(connTripNodes, getNewHopTotal(hopsTravelled), maxHops)
		}
	}

	return result
}

func getNewHopTotal(currentHops int) int {
	return currentHops + singleHop
}
