package trains

// AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
func getTestNetworkOfTrains() (*Network, map[string]*Station) {
	n, stations := NewNetwork(), NewStations("A", "B", "C", "D", "E")

	stations["A"].AddConnection(stations["B"], 5)
	stations["B"].AddConnection(stations["C"], 4)
	stations["C"].AddConnection(stations["D"], 8)
	stations["D"].AddConnection(stations["C"], 8)
	stations["D"].AddConnection(stations["E"], 6)
	stations["A"].AddConnection(stations["D"], 5)
	stations["C"].AddConnection(stations["E"], 2)
	stations["E"].AddConnection(stations["B"], 3)
	stations["A"].AddConnection(stations["E"], 7)

	for _, s := range stations {
		n.AddStation(s)
	}

	return &n, stations
}

// AB3, BC5, BD3, BE1, EC1, CF7
func getSimpleTestNetworkOfTrains() (*Network, map[string]*Station) {
	n, stations := NewNetwork(), NewStations("A", "B", "C", "D", "E", "F")

	stations["A"].AddConnection(stations["B"], 3)
	stations["B"].AddConnection(stations["C"], 5)
	stations["B"].AddConnection(stations["D"], 3)
	stations["B"].AddConnection(stations["E"], 1)
	stations["E"].AddConnection(stations["C"], 1)
	stations["C"].AddConnection(stations["F"], 7)

	for _, station := range stations {
		n.AddStation(station)
	}

	return &n, stations
}

// AB1, AC4, BC2, BD5, CA3
func getSimplerTestNetworkOfTrains() (*Network, map[string]*Station) {
	n, stations := NewNetwork(), NewStations("A", "B", "C", "D")

	stations["A"].AddConnection(stations["B"], 1)
	stations["A"].AddConnection(stations["C"], 4)
	stations["B"].AddConnection(stations["C"], 2)
	stations["B"].AddConnection(stations["D"], 5)
	stations["C"].AddConnection(stations["A"], 3)

	for _, station := range stations {
		n.AddStation(station)
	}

	return &n, stations
}
