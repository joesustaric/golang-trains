package trains

// This my representation of infinity for lack of a better way.
const INFINITY = 99999

// ShortestPathTrip this is the data structure which contains all the data
// needed to calculate the shortest path trip.
type ShortestPathTrip struct {
	myNetwork           *Network
	originalTrip        *trip
	currentNode         *station
	visitedStationTimes map[*station]int
	distanceToStation   map[*station]int
	completed           bool
}

// NewShortestPathTrip returns a newly initalised object
func NewShortestPathTrip(n *Network, t *trip) *ShortestPathTrip {
	result := &ShortestPathTrip{myNetwork: n, originalTrip: t, currentNode: t.from, completed: false}
	result.initaliseTrip()
	return result
}

func (s *ShortestPathTrip) initaliseTrip() {
	s.visitedStationTimes, s.distanceToStation = make(map[*station]int), make(map[*station]int)
	s.currentNode = s.originalTrip.from
	s.visitedStationTimes[s.currentNode] = 1
}

// CalculateConnectionsDistanceFromCurrent blah
func (s *ShortestPathTrip) CalculateConnectionsDistanceFromCurrent() {
	visitedTimes, ok := s.visitedStationTimes[s.currentNode]

	if ok && visitedTimes == 1 {
		dist, okay := s.distanceToStation[s.currentNode]

		if okay {
			for conn := range s.currentNode.connections {
				s.distanceToStation[conn] = s.currentNode.GetDistanceTo(conn) + dist
			}
		} else {
			// If we get here it should be the first trip form the origin
			for conn := range s.currentNode.connections {
				s.distanceToStation[conn] = s.currentNode.GetDistanceTo(conn)
			}
		}
	}

}

// VisitNextNode blah blah
func (s *ShortestPathTrip) VisitNextNode() {
	if !s.completed {

		shortestDist := INFINITY
		nextNode := &station{}

		for conn := range s.currentNode.connections {
			_, visited := s.visitedStationTimes[conn]
			if !visited {
				if s.currentNode.GetDistanceTo(conn) <= shortestDist {
					shortestDist = s.currentNode.GetDistanceTo(conn)
					nextNode = conn
				}
			}
		}

		s.visitedStationTimes[nextNode] = 1
		s.currentNode = nextNode

		if nextNode == s.originalTrip.to {
			s.completed = true
		}
	}
}
