package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct{
		nums []int
		rlt  int
	}{
		{
			[]int{1, 9, 5, 9, 3},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{5, 3, 2, 1},
			1,
		},
	}
	for _, tt := range tests {
		if actual := lengthOfLIS1(tt.nums); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkLengthOfLIS1(b *testing.B) {
	test := struct {
		nums []int
		rlt  int
	}{
		[]int{1, 9, 5, 9, 3},
		3,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := lengthOfLIS1(test.nums); actual != test.rlt {
			b.Errorf("Benchmark algorithm is wrong.")
		}
	}
}

func BenchmarkLengthOfLIS2(b *testing.B) {
	test := struct {
		nums []int
		rlt  int
	}{
		[]int{1, 9, 5, 9, 3},
		3,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := lengthOfLIS2(test.nums); actual != test.rlt {
			b.Errorf("Benchmark algorithm is wrong.")
		}
	}
}