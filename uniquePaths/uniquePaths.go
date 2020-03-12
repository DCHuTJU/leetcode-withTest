package main

import "fmt"

func uniquePaths(m, n int) int {
	dp := make([]int, n)
	for i:=0; i<n; i++ {
		dp[i] = 1
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func uniquePathsDP(m, n int) int {
	dp := make([][]int, n)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int, m)
	}
	for i:=0; i<m; i++ {
		dp[0][i] = 1
	}
	for i:=0; i<n; i++ {
		dp[i][0] = 1
	}
	for i:=1; i<n; i++ {
		for j:=1; j<m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

func main() {
	m, n := 7, 3
	rlt := uniquePathsDP(m, n)
	fmt.Println("Result is:", rlt)
	rlt = uniquePaths(m, n)
	fmt.Println("Result is:", rlt)
}
