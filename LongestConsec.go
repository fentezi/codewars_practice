package main

import (
	"fmt"
)

func longestConsec(strarr []string, k int) string {
	n := len(strarr)

	// Check for edge cases
	if n == 0 || k > n || k <= 0 {
		return ""
	}

	// Initialize variables to store the result
	var result string
	maxLength := 0

	// Iterate through the array
	for i := 0; i <= n-k; i++ {
		// Concatenate k consecutive strings
		currentConcatenation := ""
		for j := i; j < i+k; j++ {
			currentConcatenation += strarr[j]
		}

		// Update result if the current concatenation is longer
		if len(currentConcatenation) > maxLength {
			result = currentConcatenation
			maxLength = len(currentConcatenation)
		}
	}

	return result
}

func main() {
	// Example usage
	strarr := []string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}
	k := 2
	result := longestConsec(strarr, k)
	fmt.Println(result) // Output: folingtrashy
}
