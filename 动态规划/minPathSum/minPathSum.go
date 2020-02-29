package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i, arr := range grid {
		dp[i] = make([]int, len(arr))
	}
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	dp[0][0] = grid[0][0]
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if i==0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	rlt := minPathSum(grid)
	fmt.Println(rlt)
}