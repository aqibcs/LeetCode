package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    // Initialize a slice to store the results
    var result [][]int

    // Sort the input slice in ascending order
    sort.Ints(nums)

    // Iterate over the slice, but stop 2 elements before the end
    for i := 0; i < len(nums)-2; i++ {
        // If the current element is the same as the previous one, skip it to avoid duplicates
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // Initialize two pointers, one at the element after i, and one at the end of the slice
        left, right := i+1, len(nums)-1

        // While the left pointer is less than the right pointer
        for left < right {
            // Calculate the sum of the elements at the current positions of i, left, and right
            sum := nums[i] + nums[left] + nums[right]

            // If the sum is zero
            if sum == 0 {
                // Add the triplet to the result slice
                result = append(result, []int{nums[i], nums[left], nums[right]})

                // Skip duplicates for left
                for left < right && nums[left] == nums[left+1] {
                    left++
                }

                // Skip duplicates for right
                for left < right && nums[right] == nums[right-1] {
                    right--
                }

                // Move the left and right pointers towards each other
                left++
                right--
            } else if sum < 0 {
                // If the sum is less than zero, move the left pointer to the right
                left++
            } else {
                // If the sum is greater than zero, move the right pointer to the left
                right--
            }
        }
    }

    // Return the result slice
    return result
}

func main() {
    // Define three slices of integers
    nums1 := []int{-1,0,1,2,-1,-4}
    nums2 := []int{0,1,1}
    nums3 := []int{0,0,0}

    // Call the threeSum function with nums1 as the argument and print the result
    fmt.Println(threeSum(nums1))

    // Call the threeSum function with nums2 as the argument and print the result
    fmt.Println(threeSum(nums2))

    // Call the threeSum function with nums3 as the argument and print the result
    fmt.Println(threeSum(nums3))
}