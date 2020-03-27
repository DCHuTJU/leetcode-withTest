package main

import "testing"

func BenchmarkHasGroupSizeX(b *testing.B) {
	var test = struct {
		slice []int
		rlt   bool
	}{
		[]int{1,1,1,1,2,2,2,2,2,2},
		true,
	}
	for i:=0; i<b.N; i++ {
		if actual := hasGroupsSizeX(test.slice); actual != test.rlt {
			b.Error("Benchmark algorithm wrong.")
		}
	}
}
