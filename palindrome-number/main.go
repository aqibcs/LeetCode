package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// Convert the integer to a string
	strX := strconv.Itoa(x)

	// Compare the string with its reverse
	return strX == reverseString(strX)
}

func reverseString(s string) string {
	// Convert the string to a slice of runes and reverse it
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
