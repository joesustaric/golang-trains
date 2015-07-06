package trains

// CalculateShortestTrip blah
func CalculateShortestTrip(spt *ShortestPathTrip) {
	for spt.Completed == false {
		spt.CalcDistToConnectionsAndVisitNext()
	}
}
