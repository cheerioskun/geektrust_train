package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var destinationListA, destinationListB []string
	// Input start.
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		destinationList := strings.Fields(line)
		// First element will be of type "TRAIN_A" or TRAIN_B
		switch destinationList[0] {
		case TRAIN_A_IDENTIFIER:
			{
				// destinationList of form [TRAIN_A, ENGINE, stations...]
				destinationListA = append(destinationListA, destinationList[2:]...)
			}
		case TRAIN_B_IDENTIFIER:
			{
				destinationListB = append(destinationListB, destinationList[2:]...)
			}
		default:
			{
				panic(errors.New("INVALID INPUT"))
			}
		}
	}
	// Input complete.
	// Phase 1: Go to Hyderabad
	destinationListA = RemoveTillHyb(destinationListA, TRAIN_A_IDENTIFIER)
	destinationListB = RemoveTillHyb(destinationListB, TRAIN_B_IDENTIFIER)
	// Print the arrival orders
	fmt.Printf("ARRIVAL TRAIN_A ENGINE %s\n", strings.Join(destinationListA, " "))
	fmt.Printf("ARRIVAL TRAIN_B ENGINE %s\n", strings.Join(destinationListB, " "))

	// Phase 2: Merge
	departureList := MergeAtHyb(destinationListA, destinationListB)
	// Print departure order
	fmt.Printf("DEPARTURE TRAIN_AB ENGINE ENGINE %s\n", strings.Join(departureList, " "))
	// Done
}
