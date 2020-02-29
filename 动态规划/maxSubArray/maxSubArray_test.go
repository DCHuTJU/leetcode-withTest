package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := struct {
		nums []int
		rlt  int
	}{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, -4},
		6,
	}

	if actual := maxSubArray(tests.nums); actual != tests.rlt {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	tests := struct {
		nums []int
		rlt  int
	}{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, -4},
		6,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := maxSubArray(tests.nums); actual != tests.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
