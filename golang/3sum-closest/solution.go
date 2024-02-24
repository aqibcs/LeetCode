package main

import (
	"fmt"
	"sort"
	"math"
)

// threeSumClosest finds three integers in nums such that the sum is closest to target.
func threeSumClosest(nums []int, target int) int {
	// Sort the input array
	sort.Ints(nums)
	n := len(nums)
	// Initialize closestSum with the maximum integer value
	closestSum := math.MaxInt32

	// Iterate over the array
	for i := 0; i < n-2; i++ {
		// Initialize two pointers, one at the next position, and the other at the end of the array
		left, right := i+1, n-1

		// While the left pointer is less than the right pointer
		for left < right {
			// Calculate the sum of the values at the current position, left pointer, and right pointer
			currentSum := nums[i] + nums[left] + nums[right]

			// If the absolute difference between currentSum and target is less than the absolute difference between closestSum and target
			if abs(currentSum-target) < abs(closestSum-target) {
				// Update closestSum
				closestSum = currentSum
			}

			// If currentSum is less than target, move the left pointer to the right
			if currentSum < target {
				left++
			// If currentSum is greater than target, move the right pointer to the left
			} else if currentSum > target {
				right--
			// If currentSum is equal to target, return closestSum
			} else {
				return closestSum
			}
		}
	}

	// Return the closest sum found
	return closestSum
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example usage:
	nums1 := []int{-1, 2, 1, -4}
	target1 := 1
	// Print the sum of three numbers in nums1 that is closest to target1
	fmt.Println(threeSumClosest(nums1, target1)) // Output: 2

	nums2 := []int{0, 0, 0}
	target2 := 1
	// Print the sum of three numbers in nums2 that is closest to target2
	fmt.Println(threeSumClosest(nums2, target2)) // Output: 0
}