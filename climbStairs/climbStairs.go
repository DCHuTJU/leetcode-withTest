package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i:=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	var n = 3
	rlt := climbStairs(n)
	fmt.Println(rlt)
}

// 1298777728820984005
