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

func newTrip(from, to *Station) trip {
	return trip{from: from, to: to}
}

const (
	firstHop  int = 1
	singleHop int = 1
)

// GetNumberOfRoutes returns the number of routes in a network with a hops query.
func GetNumberOfRoutes(n *Network, thq tripHopsQuery) int {
	if origin, ok := n.GetNode(thq.origin); !ok {
		return 0
	} else if dest, ok := n.GetNode(thq.destination); !ok {
		return 0
	} else {
		return numberOfRoutes(newTrip(origin, dest), firstHop, thq.maxHops)
	}
}

func numberOfRoutes(t trip, hopsTravelled, maxHops int) int {
	result := 0
	for conn := range t.from.connections {
		if conn == t.to || result < maxHops {
			result++
		} else if hopsTravelled < maxHops {
			result += numberOfRoutes(newTrip(conn, t.to), getNewHopTotal(hopsTravelled), maxHops)
		}
	}
	return result
}

func getNewHopTotal(currentHops int) int {
	return currentHops + singleHop
}
