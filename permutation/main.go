package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums)-1 {
			currentPermutation := make([]int, len(nums))
			copy(currentPermutation, nums)
			result = append(result, currentPermutation)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	
	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(permute(nums1))

	nums2 := []int{0, 1}
	fmt.Println(permute(nums2))

	nums3 := []int{1}
	fmt.Println(permute(nums3))
}
