package main

import "fmt"

var digitMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombination(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	backtrack(&result, "", digits, 0)
	return result
}

func backtrack(result *[]string, current string, digits string, index int) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	digit := digits[index]
	letters := digitMap[digit]

	for j := 0; j < len(letters); j++ {
		backtrack(result, current+string(letters[j]), digits, index+1)
	}
}
func main() {
	fmt.Println(letterCombination("23"))
	fmt.Println(letterCombination(""))
	fmt.Println(letterCombination("2"))
}
