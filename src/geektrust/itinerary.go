package main

var (
	TRAIN_A_ITINERARY = []string{"CHN", "SLM", "BLR", "KRN", "HYB", "NGP", "ITJ", "BPL", "AGA", "NDL"}
	TRAIN_B_ITINERARY = []string{"TVC", "SRR", "MAQ", "MAO", "PNE", "HYB", "NGP", "ITJ", "BPL", "PTA", "NGP", "GHY"}
	orderA, orderB    map[string]int
)

func InitializeStationOrders() {
	for idx, val := range TRAIN_A_ITINERARY {
		orderA[val] = idx
	}
	for idx, val := range TRAIN_B_ITINERARY {
		orderB[val] = idx
	}
}
