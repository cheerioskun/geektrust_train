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
	var bogieListA, bogieListB []string
	// Input start.
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		bogieList := strings.Fields(line)
		// First element will be of type "TRAIN_A" or TRAIN_B
		switch bogieList[0] {
		case TRAIN_A_IDENTIFIER:
			{
				// bogieList of form [TRAIN_A, ENGINE, stations...]
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
	// Input complete.
	// Phase 1: Go to Hyderabad
	bogieListA = RemoveTillHyb(bogieListA, TRAIN_A_IDENTIFIER)
	bogieListB = RemoveTillHyb(bogieListB, TRAIN_B_IDENTIFIER)
	// Print the arrival orders
	fmt.Printf("ARRIVAL TRAIN_A ENGINE %s\n", strings.Join(bogieListA, " "))
	fmt.Printf("ARRIVAL TRAIN_B ENGINE %s\n", strings.Join(bogieListB, " "))

	// Phase 2: Merge
	departureList := MergeAtHyb(bogieListA, bogieListB)
	// Print departure order
	fmt.Printf("DEPARTURE TRAIN_AB ENGINE ENGINE %s\n", strings.Join(departureList, " "))
	// Done
}
