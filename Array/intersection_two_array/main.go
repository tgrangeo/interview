package main

import (
	"fmt"
	"slices"
)

func intersectionTwoArrays(nums1 []int, nums2 []int) []int {
	res := []int{}
	slices.Sort(nums1)
	slices.Sort(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 1, 1, 1, 1, 2}
	nums2 := []int{3, 1, 1, 2}
	fmt.Println(intersectionTwoArrays(nums1, nums2))
}
