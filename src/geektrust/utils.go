package main

import "sort"

const (
	HYDERABAD_STATION_STRING = "HYB"
	TRAIN_A_IDENTIFIER       = "TRAIN_A"
	TRAIN_B_IDENTIFIER       = "TRAIN_B"
)

// Removes the bogies that have been detached before arriving at Hyderabad.
// Returns the list of remaining bogies in order
func RemoveTillHyb(bogieList []string, train string) []string {
	var distanceMap map[string]int
	if train == TRAIN_A_IDENTIFIER {
		distanceMap = orderA
	} else {
		distanceMap = orderB
	}
	var finalList []string
	// For each car(characterized by its destination) check if it needs to be detached
	for _, destination := range bogieList {
		// Either destination does not belong to this itinerary or arrives before HYB
		if _, ok := distanceMap[destination]; !ok || distanceMap[destination] >= distanceMap[HYDERABAD_STATION_STRING] {
			finalList = append(finalList, destination)
		}
	}
	return finalList
}

// Merges the two trains at hyderabad and returns departure order
func MergeAtHyb(bogieListA, bogieListB []string) []string {

	// Remove all HYB bogies
	bogieListA = RemoveBogies(bogieListA, HYDERABAD_STATION_STRING)
	bogieListB = RemoveBogies(bogieListB, HYDERABAD_STATION_STRING)

	// Sort the individual trains' bogies based on distance
	sort.Slice(bogieListA, func(i, j int) bool {
		a, b := bogieListA[i], bogieListA[j]
		return distanceFromHyb[b] < distanceFromHyb[a]
	})
	sort.Slice(bogieListB, func(i, j int) bool {
		a, b := bogieListB[i], bogieListB[j]
		return distanceFromHyb[b] < distanceFromHyb[a]
	})

	var finalList []string
	// Final count of bogies is
	n := len(bogieListA) + len(bogieListB)
	// Count of how many I've joined from each
	i, j := 0, 0
	var nextBogie string
	for i+j < n {
		if i < len(bogieListA) && j < len(bogieListB) {
			// Both are available
			a, b := bogieListA[i], bogieListB[j]
			if distanceFromHyb[a] > distanceFromHyb[b] {
				// the one at the top of A is further so add it first
				nextBogie = a
				i += 1
			} else {
				nextBogie = b
				j += 1
			}
		} else if i == len(bogieListA) {
			// Only B is non empty
			nextBogie = bogieListB[j]
			j += 1
		} else {
			// Only A is non empty
			nextBogie = bogieListA[i]
			i += 1
		}
		finalList = append(finalList, nextBogie)
	}
	return finalList
}

// Remove bogies with a given destination
func RemoveBogies(bogieList []string, destination string) []string {
	i := 0
	for _, val := range bogieList {
		if val != destination {
			bogieList[i] = val
			i += 1
		}
	}
	return bogieList[:i]
}
