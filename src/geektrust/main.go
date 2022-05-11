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
	var stationListA, stationListB []string
	// Input start.
	for scanner.Scan() {
		line := scanner.Text()
		stationList := strings.Fields(line)
		// First element will be of type "TRAIN_A" or TRAIN_B
		switch stationList[0] {
		case "TRAIN_A":
			{
				// stationList of form [TRAIN_A, ENGINE, stations...]
				stationListA = append(stationListA, stationList[2:]...)
			}
		case "TRAIN_B":
			{
				stationListB = append(stationListB, stationList[2:]...)
			}
		default:
			{
				panic(errors.New("INVALID INPUT"))
			}
		}
	}
	// Input complete.

}
