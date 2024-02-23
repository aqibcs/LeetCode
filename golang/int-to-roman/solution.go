package main

import (
	"fmt"
)

func intToRoman(num int) string {
	// Define the mapping of integer values to Roman numeral symbols
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Initialize result string
	result := ""

	// Iterate through the values
	for i := 0; i < len(values); i++ {
		// Repeat the current symbol while the number is greater than or equal to the current value
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}

	return result
}

func main() {
	// Examples
	num1 := 3
	fmt.Println(intToRoman(num1)) // Expected output: "III"

	num2 := 58
	fmt.Println(intToRoman(num2)) // Expected output: "LVIII"

	num3 := 1994
	fmt.Println(intToRoman(num3)) // Expected output: "MCMXCIV"
}
