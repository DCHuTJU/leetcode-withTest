package main

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// 构建 dp 数组
	// 本题目中存在递推关系，即dp[i]是与dp[i-1]有关的，因此考虑使用动态规划
	// 当 dp[i-1]<0 时，dp[i]的值应为当前 nums[i] 值
	// 当 dp[i-1]>0 时，dp[i]的值应为 dp[i-1] + nums[i]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i:=1; i<len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	// 数据遍历后，dp 数组中的最大值即为结果
	result := -1 << 31
	for _, k := range dp {
		result = max(result, k)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func main() {
//	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, -4}
//	rlt := maxSubArray(nums)
//	fmt.Println("Result is: ", rlt)
//}
