package main

import "fmt"

func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	nums[1] = max(nums[0], nums[1])
	for i:=2; i<len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var nums = []int{1, 2, 3, 1}
	rlt := rob(nums)
	fmt.Println(rlt)
}