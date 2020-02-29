package main

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		rlt 	 int
	} {
		{
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
	}
	for _, tt := range tests {
		if actual := minimumTotal(tt.triangle); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkMinimumTotal(b *testing.B) {
	test := struct {
		triangle [][]int
		rlt 	 int
	} {
		[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		11,
	}

	b.ResetTimer()
	for i:=1; i<b.N; i++ {
		if actual := minimumTotal(test.triangle); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
