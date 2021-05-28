package main

import "testing"

func BenchmarkHammingDistance(b *testing.B) {
	x, y := 1, 4
	for i:=0; i<b.N; i++ {
		if rlt := hammingDistance(x, y); rlt != 2 {
			b.Error("Benchmark algorithm wrong.")
		}
	}
}

func BenchmarkTotalHammingDistance(b *testing.B) {
	input := []int{4, 14, 2}
	for i:=0; i<b.N; i++ {
		if rlt := totalHammingDistance(input); rlt != 6 {
			b.Error("Benchmark algorithm wrong.")
		}
	}
}
