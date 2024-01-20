package main

import (
	"fmt"
)

func maxArea(height []int) int {
    // Initialize pointers
    left := 0
    right := len(height) - 1

    // Initialize variable to store maximum area
    maxArea := 0

    for left < right {
        // Calculate the width (distance between lines)
        width := right - left

        // Calculate the height (minimum of the two lines)
        h := min(height[left], height[right])

        // Update maxArea if the current area is greater
        maxArea = max(maxArea, width*h)

        // Move pointers
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

func main() {
    height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println(maxArea(height1))  // Expected output: 49

    height2 := []int{1, 1}
    fmt.Println(maxArea(height2))  // Expected output: 1
}
