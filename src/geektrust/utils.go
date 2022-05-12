package main

import (
	"bufio"
	"errors"
	"sort"
	"strings"
)

const (
	HYDERABAD_STATION_STRING = "HYB"
	TRAIN_A_IDENTIFIER       = "TRAIN_A"
	TRAIN_B_IDENTIFIER       = "TRAIN_B"
)

// Reads input from a scanner and parses into two lists
func ParseInput(scanner *bufio.Scanner) ([]string, []string) {
	var bogieListA, bogieListB []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		bogieList := strings.Fields(line)

		switch bogieList[0] {
		case TRAIN_A_IDENTIFIER:
			{
				bogieListA = append(bogieListA, bogieList[2:]...)
			}
		case TRAIN_B_IDENTIFIER:
			{
				bogieListB = append(bogieListB, bogieList[2:]...)
			}
		default:
			{
				panic(errors.New("INVALID INPUT"))
			}
		}
	}
	return bogieListA, bogieListB
}

// Removes the bogies that have been detached before arriving at Hyderabad and returns the list of remaining bogies in order
func RemoveTillHyderabad(bogieList []string, train string) []string {
	var distanceMap map[string]int
	if train == TRAIN_A_IDENTIFIER {
		distanceMap = orderA
	} else {
		distanceMap = orderB
	}
	var finalList []string
	for _, destination := range bogieList {
		// Either destination does not belong to this itinerary or arrives before HYB
		if _, ok := distanceMap[destination]; !ok || distanceMap[destination] >= distanceMap[HYDERABAD_STATION_STRING] {
			finalList = append(finalList, destination)
		}
	}
	return finalList
}

// Merges the two trains at hyderabad and returns departure order
func MergeAtHyderabad(bogieListA, bogieListB []string) []string {
	// Remove all Hyderabad bogies
	bogieListA = RemoveBogies(bogieListA, HYDERABAD_STATION_STRING)
	bogieListB = RemoveBogies(bogieListB, HYDERABAD_STATION_STRING)

	// Sort the individual trains' bogies based on distance
	sort.Slice(bogieListA, func(i, j int) bool {
		first, second := bogieListA[i], bogieListA[j]
		return distanceFromHyderabad[second] < distanceFromHyderabad[first]
	})
	sort.Slice(bogieListB, func(i, j int) bool {
		first, second := bogieListB[i], bogieListB[j]
		return distanceFromHyderabad[second] < distanceFromHyderabad[first]
	})

	finalList := MergeSorted(bogieListA, bogieListB)
	return finalList
}

// Merge two sorted lists
func MergeSorted(bogieListA, bogieListB []string) []string {
	var finalList []string
	finalTrainLength := len(bogieListA) + len(bogieListB)
	indexA, indexB := 0, 0
	var nextBogie string
	for indexA+indexB < finalTrainLength {
		if indexA < len(bogieListA) && indexB < len(bogieListB) {
			// Both are available
			stationA, stationB := bogieListA[indexA], bogieListB[indexB]
			if distanceFromHyderabad[stationA] > distanceFromHyderabad[stationB] {
				nextBogie = stationA
				indexA += 1
			} else {
				nextBogie = stationB
				indexB += 1
			}
		} else if indexA == len(bogieListA) {
			// Only B is non empty
			nextBogie = bogieListB[indexB]
			indexB += 1
		} else {
			// Only A is non empty
			nextBogie = bogieListA[indexA]
			indexA += 1
		}
		finalList = append(finalList, nextBogie)
	}
	return finalList
}

// Remove bogies with a given destination
func RemoveBogies(bogieList []string, destination string) []string {
	nextInsertPosition := 0
	for _, val := range bogieList {
		if val != destination {
			bogieList[nextInsertPosition] = val
			nextInsertPosition += 1
		}
	}
	return bogieList[:nextInsertPosition]
}
