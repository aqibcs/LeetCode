package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		// Get the last digit of x
		pop := x % 10
		x /= 10

		// Check for integer overflow before updating the result
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && pop < -8) {
			return 0
		}

		// Update the result by appending the last digit
		result = result*10 + pop
	}

	return result
}

func main() {
	// Example usage
	input1 := 123
	output1 := reverse(input1)
	fmt.Printf("Input: %d, Output: %d\n", input1, output1)

	input2 := -123
	output2 := reverse(input2)
	fmt.Printf("Input: %d, Output: %d\n", input2, output2)

	input3 := 120
	output3 := reverse(input3)
	fmt.Printf("Input: %d, Output: %d\n", input3, output3)
}
