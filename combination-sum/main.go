package main

import "fmt"

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(&result, candidates, []int{}, target, 0)
	return result
}

func backtrack(result *[][]int, candidates []int, currentCombination []int, remainingTarget, start int) {
	if remainingTarget == 0 {
		*result = append(*result, append([]int{}, currentCombination...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > remainingTarget {
			continue
		}

		currentCombination = append(currentCombination, candidates[i])

		backtrack(result, candidates, currentCombination, remainingTarget-candidates[i], i)

		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func main() {
	fmt.Println(CombinationSum([]int{2, 3, 6, 7}, 7))

	fmt.Println(CombinationSum([]int{2, 3, 5}, 8))

	fmt.Println(CombinationSum([]int{2}, 1))
}
