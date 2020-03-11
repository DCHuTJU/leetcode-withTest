package main

import (
	"fmt"
	"github.com/dengchengH/go-simple-func/compare"
	"math"
)

func integerBreak(n int) int {
	if n <= 3 {
		return n-1;
	}
	x, y := n/3, n%3
	if y == 0 {
		return int(math.Pow(3, float64(x)))
	}
	if y == 1 {
		return int(math.Pow(3, float64(x-1))*4)
	}
	return int(math.Pow(3, float64(x))*2)
}

func integerBreakDP(n int) int {
	dp := make([]int, n+1)
	// 定义状态
	dp[1] = 1
	for i:=2; i<=n; i++ {
		for j:=1; j<i; j++ {
			dp[i] = compare.Max(dp[i], compare.Max(dp[j], j) * (i - j))
		}
	}
	return dp[n]
}

func main() {
	var a = 10
	rlt := integerBreak(a)
	fmt.Println("Result is: ", rlt)
	rlt = integerBreakDP(a)
	fmt.Println("Result of DP is: ", rlt)
}