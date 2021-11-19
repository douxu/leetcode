package main

import (
	"fmt"
	algorithms "leetcode/Algorithms"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(algorithms.FindMedianSortedArrays(nums1, nums2))
	fmt.Println(algorithms.FindMedianSortedArrays1(nums1, nums2))
}
