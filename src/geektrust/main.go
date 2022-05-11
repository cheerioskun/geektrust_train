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
	var dstListA, dstListB []string
	// Input start.
	for scanner.Scan() {
		line := scanner.Text()
		dstList := strings.Fields(line)
		// First element will be of type "TRAIN_A" or TRAIN_B
		switch dstList[0] {
		case "TRAIN_A":
			{
				// dstList of form [TRAIN_A, ENGINE, stations...]
				dstListA = append(dstListA, dstList[2:]...)
			}
		case "TRAIN_B":
			{
				dstListB = append(dstListB, dstList[2:]...)
			}
		default:
			{
				panic(errors.New("INVALID INPUT"))
			}
		}
	}
	// Input complete.
	// Phase 1: Go to Hyderabad
	dstListA = ArriveAtHyb(dstListA, "A")
	dstListB = ArriveAtHyb(dstListB, "B")
	// Print the arrival orders
	fmt.Printf("ARRIVAL TRAIN_A ENGINE %s\n", strings.Join(dstListA, " "))
	fmt.Printf("ARRIVAL TRAIN_B ENGINE %s\n", strings.Join(dstListB, " "))

}
