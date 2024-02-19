package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		nums1 = append(nums1, nums2[i])
	}

	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	size := len(nums1)
	if size%2 == 0 {
		n1 := float64(nums1[size/2-1])
		n2 := float64(nums1[size/2])
		return (n1 + n2) / 2
	} else {
		return float64(nums1[size/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
