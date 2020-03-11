package main

import "testing"

func BenchmarkIsStraightSort(b *testing.B) {
	var test = struct {
		vals []int
		rlt  bool
	}{
		[]int{0, 1, 2, 5, 3},
		true,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := isStraightSort(test.vals); actual != test.rlt {
			b.Errorf("Benchmark Algorithm is wrong.")
		}
	}
}

func BenchmarkIsStraightWithoutSort(b *testing.B) {
	var test = struct {
		vals []int
		rlt  bool
	}{
		[]int{0, 1, 2, 5, 3},
		true,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := isStraightWithoutSort(test.vals); actual != test.rlt {
			b.Errorf("Benchmark Algorithm is wrong.")
		}
	}
}
