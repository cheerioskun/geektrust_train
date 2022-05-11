package main

import "testing"

func TestRemoveTillHyb(t *testing.T) {
	destinationListA := []string{"HYB", "BLR", "SLM", "HYB", "ITJ", "NDL"}
	processedListA := RemoveTillHyb(destinationListA, TRAIN_A_IDENTIFIER)
	if !isEqualSlice(processedListA, []string{"HYB", "HYB", "ITJ", "NDL"}) {
		t.Fatalf("Failed on case: TRAIN_A: %v\n", destinationListA)
	}
	destinationListB := []string{"MAO", "NGP", "ITJ", "TVC", "MAQ", "SRR"}
	processedListB := RemoveTillHyb(destinationListB, TRAIN_B_IDENTIFIER)
	if !isEqualSlice(processedListB, []string{"NGP", "ITJ"}) {
		t.Fatalf("Failed on case: TRAIN_B: %v\n", destinationListB)
	}
}

func TestMergeAtHyb(t *testing.T) {
	destinationListA := []string{"HYB", "HYB", "ITJ", "NDL"}
	destinationListB := []string{"GHY", "NGP", "ITJ"}
	departureList := MergeAtHyb(destinationListA, destinationListB)
	if !isEqualSlice(departureList, []string{"GHY", "NDL", "ITJ", "ITJ", "NGP"}) {
		t.Fatalf("Failed on case: TRAIN_A: %v, TRAIN_B: %v", destinationListA, destinationListB)
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
