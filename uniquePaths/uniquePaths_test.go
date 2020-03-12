package main

import "testing"

func BenchmarkUniquePathsDP(b *testing.B) {
	test := struct {
		m   int
		n   int
		rlt int
	}{
		7,
		3,
		28,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual:=uniquePathsDP(test.m, test.n); actual!=test.rlt {
			b.Error("Benchmark sample is wrong.")
		}
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	test := struct {
		m   int
		n   int
		rlt int
	}{
		7,
		3,
		28,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual:=uniquePaths(test.m, test.n); actual!=test.rlt {
			b.Error("Benchmark sample is wrong.")
		}
	}
}
