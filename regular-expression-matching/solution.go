package main

import "fmt"

func isMatch(s string, p string) bool {
	// Initialize a 2D DP array with all false values
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	// Empty pattern matches empty string
	dp[0][0] = true

	// Fill in the first row for patterns with '*'
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill in the DP array based on the recurrence relation
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return dp[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("aa", "a"))

	fmt.Println(isMatch("aa", "a*"))

	fmt.Println(isMatch("aa", ".*"))
}
