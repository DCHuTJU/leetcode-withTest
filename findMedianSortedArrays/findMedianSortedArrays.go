package main

import (
	"github.com/dengchengH/go-simple-func/compare"
	"math"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	total := len1 + len2
	left, right := (total + 1) / 2, (total + 2) / 2
	return float64(findK(nums1, 0, nums2, 0, left) + findK(nums1, 0, nums2, 0, right)) / 2.0
}

func findK(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[j+k-1]
	}
	if k == 1 {
		return compare.Min(nums1[i], nums2[j])
	}
	// 计算出每次要比较的两个数的值，来决定 "删除"" 哪边的元素
	var mid1, mid2 int
	if j + k / 2 - 1 < len(nums1) {
		mid1 = nums1[j+k/2-1]
	} else {
		mid1 = math.MinInt32
	}
	if j + k / 2 - 1 < len(nums2) {
		mid2 = nums1[j+k/2-1]
	} else {
		mid2 = math.MinInt32
	}
	if mid1 < mid2 {
		return findK(nums1, i + k / 2, nums2, j, k - k / 2)
	}
	return findK(nums1, i, nums2, j + k / 2, k - k / 2)
}


