package main

import "sort"

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
		// Either dst does not belong to this itinerary or arrives before HYB
		if _, ok := distanceMap[dst]; !ok || distanceMap[dst] >= distanceMap["HYB"] {
			finalList = append(finalList, dst)
		}
	}
	return finalList
}

// Merges the two trains at hyderabad and returns departure order
func MergeAtHyb(dstListA, dstListB []string) []string {
	// Create a reference map to check distances of all subsequent stations from HYB
	distanceFromHyb := make(map[string]int)
	for station, distFromA := range orderA {
		// The map will have negative values for already passed stations but that is fine
		distanceFromHyb[station] = distFromA - orderA["HYB"]
	}
	for station, distFromB := range orderB {
		distanceFromHyb[station] = distFromB - orderB["HYB"]
	}

	// Remove all HYB cars now
	dstListA = RemoveCars(dstListA, "HYB")
	dstListB = RemoveCars(dstListB, "HYB")

	// Sort the individual trains' cars based on distance
	sort.Slice(dstListA, func(i, j int) bool {
		return distanceFromHyb[dstListA[j]] < distanceFromHyb[dstListA[i]]
	})
	sort.Slice(dstListB, func(i, j int) bool {
		return distanceFromHyb[dstListB[j]] < distanceFromHyb[dstListB[i]]
	})

	// Now let's merge
	var finalList []string
	// Final count of cars is
	n := len(dstListA) + len(dstListB)
	// Count of how many I've joined from each
	i, j := 0, 0
	for i+j < n {

		if i < len(dstListA) && j < len(dstListB) {
			// Both are available
			if distanceFromHyb[dstListA[i]] > distanceFromHyb[dstListB[j]] {
				// the one at the top of A is further so add it first
				finalList = append(finalList, dstListA[i])
				i += 1
			} else {
				finalList = append(finalList, dstListB[j])
				j += 1
			}
		} else if i == len(dstListA) {
			// Only B is non empty
			finalList = append(finalList, dstListB[j])
			j += 1
		} else {
			// Only A is non empty
			finalList = append(finalList, dstListA[i])
			i += 1
		}
	}
	return finalList
}

// Remove cars with a given destination
func RemoveCars(dstList []string, dst string) []string {
	i := 0
	for _, val := range dstList {
		if val != dst {
			dstList[i] = val
			i += 1
		}
	}
	return dstList[:i]
}
