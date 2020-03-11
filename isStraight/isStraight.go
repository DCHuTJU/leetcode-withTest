package main

import (
	"fmt"
	"sort"
)

func isStraightSort(nums []int) bool {
	sort.Ints(nums)
	// 增加一个额外空间
	sub := 0
	for i:=0; i<4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
		// 有大王的情况这样就解决不了
		sub += nums[i+1] - nums[i]
	}
	return sub < 5
}

func isStraightWithoutSort(nums []int) bool {
	arr := make([]int, 14)
	for i:=0; i<len(nums); i++ {
		arr[nums[i]]++
	}
	for i:=1; i<len(arr); i++ {
		if arr[i] > 1 {
			return false
		}
	}
	min, max := 1, 13
	for min < 14 && arr[min] == 0 {
		min++
	}
	for max >=0 && arr[max] == 0 {
		max--
	}
	return (max-min) <= 4
}

func main() {
	var a = []int{0, 1, 2, 5, 3}
	rlt := isStraightSort(a)
	fmt.Println("Result is: ", rlt)
	rlt = isStraightWithoutSort(a)
	fmt.Println("Result is: ", rlt)
}
