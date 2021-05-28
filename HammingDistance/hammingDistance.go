package main

// 计算汉明距离以及相关例题
// Leetcode-461
// 简单题

func hammingDistance(x, y int) int {
	// 1.使用自带库
	// return bits.OnesCount(uint(x ^ y))
	// 2.移位实现计数
	// 记 s = x^y 统计 1 出现的次数
	ans := 0
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return ans
}

// Leetcode-477
// 中等题

func totalHammingDistance(nums []int) int {
	ans := 0
	n := len(nums)
	for i:=0; i<30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		ans += c * (n-c)
	}
	return ans
}
