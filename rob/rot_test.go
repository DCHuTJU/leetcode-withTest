package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		rlt  int
	} {
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{2, 7, 9, 3, 1},
			12,
		},
	}
	for _, tt := range tests {
		if actual := rob(tt.nums); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkRob(b *testing.B) {
	test := struct {
		nums []int
		rlt  int
	}{
		[]int{2, 7, 9, 3, 1},
		12,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := rob(test.nums); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong. And result is %d.", actual)
		}
	}
}
