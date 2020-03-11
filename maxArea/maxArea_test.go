package main

import "testing"

// 开始只更新bench测试，test测试可以自行处理
func BenchmarkMaxArea(b *testing.B) {
	var test = struct {
		vals []int
		rlt  int
	}{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := maxArea(test.vals); actual != test.rlt {
			b.Errorf("Benchmark is wrong.")
		}
	}
}
