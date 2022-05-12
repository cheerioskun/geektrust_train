package main

import (
	"bufio"
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
	bogieListA, bogieListB := ParseInput(scanner)

	bogieListA = RemoveTillHyderabad(bogieListA, TRAIN_A_IDENTIFIER)
	bogieListB = RemoveTillHyderabad(bogieListB, TRAIN_B_IDENTIFIER)

	fmt.Printf("ARRIVAL TRAIN_A ENGINE %s\n", strings.Join(bogieListA, " "))
	fmt.Printf("ARRIVAL TRAIN_B ENGINE %s\n", strings.Join(bogieListB, " "))

	departureList := MergeAtHyderabad(bogieListA, bogieListB)

	fmt.Printf("DEPARTURE TRAIN_AB ENGINE ENGINE %s\n", strings.Join(departureList, " "))

}
