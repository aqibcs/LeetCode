package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	// Step 1: Read and ignore leading whitespace
	i := 0
	for i < len(s) && unicode.IsSpace(rune(s[i])) {
		i++
	}

	// Step 2: check for '+' or '-'
	sign := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3: Read digits until a non-digit character is encountered
	result := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Step 4: check for integer overflow
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	// Step 5: Apply the sign
	return sign * result
}

func main() {
	input1 := "42"
	result1 := myAtoi(input1)
	fmt.Println(result1)

	input2 := "	-42"
	result2 := myAtoi(input2)
	fmt.Println(result2)

	input3 := "4193 with words"
	result3 := myAtoi(input3)
	fmt.Println(result3)
}
