package main

import "testing"

func BenchmarkFindSubsequences(b *testing.B) {
	nums := []int{4, 6, 7, 7}
	for i:=0; i<b.N; i++ {
		findSubsequences(nums)
	}
}
