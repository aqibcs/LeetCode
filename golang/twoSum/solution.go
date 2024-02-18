package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numIndices[complement]; found {
			return []int{index, i}
		}

		numIndices[num] = i
	}
	return nil
}

func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Println("Example 1:", result1)

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Println("Example 2:", result2)

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Println("Example 3:", result3)
}
