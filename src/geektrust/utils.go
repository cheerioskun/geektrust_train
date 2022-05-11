package main

import "sort"

const (
	HYDERABAD_STATION_STRING = "HYB"
	TRAIN_A_IDENTIFIER       = "TRAIN_A"
	TRAIN_B_IDENTIFIER       = "TRAIN_B"
)

// Removes the bogies that have been detached before arriving at Hyderabad.
// Returns the list of remaining bogies in order
func RemoveTillHyb(destinationList []string, train string) []string {
	var distanceMap map[string]int
	if train == TRAIN_A_IDENTIFIER {
		distanceMap = orderA
	} else {
		distanceMap = orderB
	}
	var finalList []string
	// For each car(characterized by its destination) check if it needs to be detached
	for _, destination := range destinationList {
		// Either destination does not belong to this itinerary or arrives before HYB
		if _, ok := distanceMap[destination]; !ok || distanceMap[destination] >= distanceMap[HYDERABAD_STATION_STRING] {
			finalList = append(finalList, destination)
		}
	}
	return finalList
}

// Merges the two trains at hyderabad and returns departure order
func MergeAtHyb(destinationListA, destinationListB []string) []string {

	distanceFromHyb := CalculateDistancesFromHyb()
	// Remove all HYB bogies now
	destinationListA = RemoveBogies(destinationListA, HYDERABAD_STATION_STRING)
	destinationListB = RemoveBogies(destinationListB, HYDERABAD_STATION_STRING)

	// Sort the individual trains' bogies based on distance
	sort.Slice(destinationListA, func(i, j int) bool {
		return distanceFromHyb[destinationListA[j]] < distanceFromHyb[destinationListA[i]]
	})
	sort.Slice(destinationListB, func(i, j int) bool {
		return distanceFromHyb[destinationListB[j]] < distanceFromHyb[destinationListB[i]]
	})

	// Now let's merge
	var finalList []string
	// Final count of bogies is
	n := len(destinationListA) + len(destinationListB)
	// Count of how many I've joined from each
	i, j := 0, 0
	var nextBogie string
	for i+j < n {
		if i < len(destinationListA) && j < len(destinationListB) {
			// Both are available
			if distanceFromHyb[destinationListA[i]] > distanceFromHyb[destinationListB[j]] {
				// the one at the top of A is further so add it first
				nextBogie = destinationListA[i]
				i += 1
			} else {
				nextBogie = destinationListB[j]
				j += 1
			}
		} else if i == len(destinationListA) {
			// Only B is non empty
			nextBogie = destinationListB[j]
			j += 1
		} else {
			// Only A is non empty
			nextBogie = destinationListA[i]
			i += 1
		}
		finalList = append(finalList, nextBogie)
	}
	return finalList
}

// Remove bogies with a given destination
func RemoveBogies(destinationList []string, destination string) []string {
	i := 0
	for _, val := range destinationList {
		if val != destination {
			destinationList[i] = val
			i += 1
		}
	}
	return destinationList[:i]
}

// Creates a reference map to check distances of all subsequent stations from HYB
func CalculateDistancesFromHyb() map[string]int {
	distanceFromHyb := make(map[string]int)
	for station, distFromA := range orderA {
		// The map will have negative values for already passed stations but that is fine
		distanceFromHyb[station] = distFromA - orderA[HYDERABAD_STATION_STRING]
	}
	for station, distFromB := range orderB {
		distanceFromHyb[station] = distFromB - orderB[HYDERABAD_STATION_STRING]
	}
	return distanceFromHyb
}
