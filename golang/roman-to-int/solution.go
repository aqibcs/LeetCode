package main

import (
	"fmt"
	"strings"
)

// romanToInt converts a Roman numeral string to an integer.
func romanToInt(s string) int {
	// Define a map to store the values of Roman numerals.
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Replace specific combinations to simplify processing.
	s = strings.Replace(s, "CM", "CCCCCCCCC", -1) // 900
	s = strings.Replace(s, "CD", "CCCC", -1)      // 400
	s = strings.Replace(s, "XC", "XXXXXXXXX", -1) // 90
	s = strings.Replace(s, "XL", "XXXX", -1)      // 40
	s = strings.Replace(s, "IX", "IIIIIIIII", -1) // 9
	s = strings.Replace(s, "IV", "IIII", -1)      // 4

	var sum int
	// Iterate through the modified string and accumulate the corresponding values.
	for _, roman := range s {
		sum += romans[roman]
	}
	return sum
}

func main() {
	// Examples
	input1 := "III"
	fmt.Println(romanToInt(input1))

	input2 := "LVIII"
	fmt.Println(romanToInt(input2))

	input3 := "MCMXCIV"
	fmt.Println(romanToInt(input3))
}
