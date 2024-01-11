package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[byte]int)

	start := 0
	result := 0

	// Iterate through the string
	for end := 0; end < len(s); end++ {
		// If the character is already in the map and its last index is after the start of the window,
		// update the start of the window
		if lastIndexPos, found := lastIndex[s[end]]; found && lastIndexPos >= start {
			start = lastIndexPos + 1
		}

		// Update the last index of the current character in the map
		lastIndex[s[end]] = end

		// Update the result if the current substring is longer
		if end-start+1 > result {
			result = end - start + 1
		}
	}

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
