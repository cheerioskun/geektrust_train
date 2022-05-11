package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	filePath := "./sample_input/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	bogieListA, bogieListB := ParseInput(scanner)
	if !isEqualSlice(bogieListA, []string{"SLM", "BLR", "KRN", "HYB", "SLM", "NGP", "ITJ"}) ||
		!isEqualSlice(bogieListB, []string{"SRR", "MAO", "NJP", "PNE", "PTA"}) {
		t.Fail()
	}
}
func TestRemoveTillHyderabad(t *testing.T) {
	bogieListA := []string{"HYB", "BLR", "SLM", "HYB", "ITJ", "NDL"}
	processedListA := RemoveTillHyderabad(bogieListA, TRAIN_A_IDENTIFIER)
	if !isEqualSlice(processedListA, []string{"HYB", "HYB", "ITJ", "NDL"}) {
		t.Fatalf("Failed on case: TRAIN_A: %v\n", bogieListA)
	}
	bogieListB := []string{"MAO", "NGP", "ITJ", "TVC", "MAQ", "SRR"}
	processedListB := RemoveTillHyderabad(bogieListB, TRAIN_B_IDENTIFIER)
	if !isEqualSlice(processedListB, []string{"NGP", "ITJ"}) {
		t.Fatalf("Failed on case: TRAIN_B: %v\n", bogieListB)
	}
}

func TestMergeAtHyderabad(t *testing.T) {
	bogieListA := []string{"HYB", "HYB", "ITJ", "NDL"}
	bogieListB := []string{"GHY", "NGP", "ITJ"}
	departureList := MergeAtHyderabad(bogieListA, bogieListB)
	if !isEqualSlice(departureList, []string{"GHY", "NDL", "ITJ", "ITJ", "NGP"}) {
		t.Fatalf("Failed on case: TRAIN_A: %v, TRAIN_B: %v", bogieListA, bogieListB)
	}
}

func TestRemoveBogies(t *testing.T) {
	bogieList := []string{"HYB", "BLR", "SLM", "HYB", "ITJ", "NDL"}
	processedList := RemoveBogies(bogieList, "HYB")
	if !isEqualSlice(processedList, []string{"BLR", "SLM", "ITJ", "NDL"}) {
		t.Fail()
	}
}

func isEqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
