package trains

type tripHopsQuery struct {
	origin      string
	destination string
	maxHops     int
}

type trip struct {
	from *Station
	to   *Station
}

const firstHop int = 1
const singleHop int = 1

// GetNumberOfRoutes returns the number of routes in a network with a hops query.
func GetNumberOfRoutes(n *Network, thq tripHopsQuery) int {
	if origin, ok := n.GetNode(thq.origin); !ok {
		return 0
	} else if dest, ok := n.GetNode(thq.destination); !ok {
		return 0
	} else {
		tripNodes := trip{from: origin, to: dest}
		return numberOfRoutes(tripNodes, firstHop, thq.maxHops)
	}
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
