package main

var (
	// Map station to distance from origin
	orderA = map[string]int{
		"CHN": 0,
		"SLM": 350,
		"BLR": 550,
		"KRN": 900,
		"HYB": 1200,
		"NGP": 1600,
		"ITJ": 1900,
		"BPL": 2000,
		"AGA": 2500,
		"NDL": 2700,
	}

	orderB = map[string]int{
		"TVC": 0,
		"SRR": 300,
		"MAQ": 600,
		"MAO": 1000,
		"PNE": 1400,
		"HYB": 2000,
		"NGP": 2400,
		"ITJ": 2700,
		"BPL": 2800,
		"PTA": 3800,
		"NJP": 4200,
		"GHY": 4700,
	}
	distanceFromHyderabad map[string]int
)

func init() {
	distanceFromHyderabad = make(map[string]int)
	for station, distFromA := range orderA {
		// The map will have negative values for already passed stations but that is fine
		distanceFromHyderabad[station] = distFromA - orderA[HYDERABAD_STATION_STRING]
	}
	for station, distFromB := range orderB {
		distanceFromHyderabad[station] = distFromB - orderB[HYDERABAD_STATION_STRING]
	}
}
