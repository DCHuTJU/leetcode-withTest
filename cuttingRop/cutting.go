package cuttingRop


func CuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1 // 首项
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k])) // 递推公式
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
