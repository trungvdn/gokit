package main

import (
	"fmt"
	"strings"
)

func findRepeatedSequences(s string, k int) string {
	windowSize := k

	if len(s) <= windowSize {
		return "The string is too short."
	} else {
		return "The string is long enough."
	}
}

// Driver code
func slidingwindow() {
	inputsString := []string{"ACGT", "AGACCTAGAC", "AAAAACCCCCAAAAACCCCCC",
		"GGGGGGGGGGGGGGGGGGGGGGGGG", "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
		"TTTTTGGGTTTTCCA",
		"", "AAAAAACCCCCCCAAAAAAAACCCCCCCTG", "ATATATATATATATAT"}
	inputsK := []int{3, 3, 8, 10, 13, 30, 40, 30, 21}

	for i := range inputsK {
		fmt.Printf("%d.\tInput Sequence: \"%s\"\n", i+1, inputsString[i])
		fmt.Printf("\tk: %d\n", inputsK[i])
		fmt.Printf("\tComparing k with the length of the input string: \"%s\"",
			findRepeatedSequences(inputsString[i], inputsK[i]))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
