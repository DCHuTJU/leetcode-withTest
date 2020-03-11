package main

import (
	"fmt"
	"github.com/dengchengH/go-simple-func/compare"
)

func maxArea(nums []int) int {
	i, j := 0, len(nums)-1
	area, rlt := 0, 0
	for i != j {
		area = compare.Min(nums[i], nums[j]) * (j - i)
		if area > rlt {
			rlt = area
		}
		if nums[i] > nums[j] {
			j--
		} else {
			i++
		}
	}
	return rlt
}

func main() {
	var a = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	rlt := maxArea(a)
	fmt.Println("Max Area is:", rlt)
}
