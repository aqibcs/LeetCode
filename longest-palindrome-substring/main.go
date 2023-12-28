package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// dp[i][j] is true if the substring s[i:j+1] is a palindrome
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	start, maxLength := 0, 1

	// Check palindromes of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Check palindromes of length 3 or more
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLength = length
			}
		}
	}

	return s[start : start+maxLength]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
