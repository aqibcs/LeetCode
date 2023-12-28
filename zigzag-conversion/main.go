package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}

	zigzag := make([]byte, len(s))

	maxJump := (numRows - 1) * 2
	k := 0
	for i := 0; i < numRows; i++ {
		j := i

		available := maxJump
		jump := maxJump - (i * 2)
		for j < len(s) && k < len(zigzag) {
			zigzag[k] = s[j]
			k++

			if maxJump == jump || jump == 0 {
				j += maxJump
			} else if available == maxJump {
				j += jump
				available -= jump
			} else {
				j += available
				available = maxJump
			}
		}

	}

	return string(zigzag)
}

func main() {
	// Example 1
	s1 := "PAYPALISHIRING"
	numRows1 := 3
	fmt.Println(convert(s1, numRows1))

	// Example 2
	s2 := "PAYPALISHIRING"
	numRows2 := 4
	fmt.Println(convert(s2, numRows2))

	// Example 3
	s3 := "A"
	numRows3 := 1
	fmt.Println(convert(s3, numRows3))
}
