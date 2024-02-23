package main

import (
	"fmt"
)

// longestCommonPrefix finds the longest common prefix string among an array of strings.
// If there is no common prefix, it returns an empty string.
func longestCommonPrefix(strs []string) string {
	// Check if the input array is empty
	if len(strs) == 0 {
		return ""
	}

	// Initialize the result with the first string in the array
	result := strs[0]

	// Iterate through the rest of the strings in the array
	for i := 1; i < len(strs); i++ {
		// Compare characters at each position
		j := 0
		for j < len(result) && j < len(strs[i]) && result[j] == strs[i][j] {
			j++
		}

		// Shorten the result string based on the mismatch
		result = result[:j]

		// If the result becomes an empty string, there is no common prefix
		if result == "" {
			return ""
		}
	}

	// Return the final result, which is the longest common prefix
	return result
}

func main() {
	// Example 1
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1))

	// Example 2
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs2))
}
