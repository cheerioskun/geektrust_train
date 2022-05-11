package main

// Removes the cars that have been detached before arriving at Hyderabad
// Returns the list of remaining cars in order
func ArriveAtHyb(dstList []string, train string) []string {
	var distanceMap map[string]int
	if train == "A" {
		distanceMap = orderA
	} else {
		distanceMap = orderB
	}
	var finalList []string
	// For each car(characterized by its destination) check if it needs to be detached
	for _, dst := range dstList {
		if distanceMap[dst] < distanceMap["HYB"] {
			finalList = append(finalList, dst)
		}
	}
	return finalList
}
