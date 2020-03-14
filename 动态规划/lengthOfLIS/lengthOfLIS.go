package main

import "fmt"

func lengthOfLIS1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i:=0; i<len(nums); i++ {
		dp[i] = 1
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	heaps := []int{nums[0]}
	for i := 1; i < len(nums); i ++ {
		pos := binarySearch(heaps,0,len(heaps)-1,nums[i])
		// 在开头，说明有更小的值，则更新当前最小值
		if pos == -1 {
			heaps[0] = nums[i]
			// 在 heap 尾部，说明有更大的值，就加入进来
		}else if pos == len(heaps) {
			heaps = append(heaps,nums[i])
		}else {
			// 若在中间位置，直接替换即可
			heaps[pos] = nums[i]
		}
	}
	// 最后长度即为最长上升子序列
	return len(heaps)
}

func binarySearch(nums []int,l,r int,target int) int {
	for {
		if l > r {
			break
		}
		mid := l + (r-l) >> 1
		if nums[mid] < target {
			l = mid+1
		}else {
			r = mid-1
		}
	}
	return l
}

func main() {
	var nums = []int{5, 3, 2, 1}
	rlt := lengthOfLIS1(nums)
	fmt.Println("Result is: ", rlt)
	rlt = lengthOfLIS2(nums)
	fmt.Println("Result is: ", rlt)
}
