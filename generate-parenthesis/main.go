package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, max int, result *[]string) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}
	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}

func main() {
	// Example 1
	fmt.Println(generateParenthesis(3))

	// Example 2
	fmt.Println(generateParenthesis(1))
}
