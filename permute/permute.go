package main

import (
	"fmt"
	"math"
)

var (
	temp []int
	ans [][]int
)

func findSubsequences(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur int, last int, nums []int) {
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}

	// 选择当前节点
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur + 1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	// 不选当前节点
	if nums[cur] != last {
		dfs(cur + 1, last, nums)
	}
}

func main() {
	nums := []int{4, 6, 7, 7}
	rlt := findSubsequences(nums)
	fmt.Println(rlt)
}
