package main

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		rlt  int
	} {
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
	}

	for _, tt := range tests {
		if actual := minPathSum(tt.grid); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	test := struct {
		grid [][]int
		rlt  int
	} {
		[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
		7,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := minPathSum(test.grid); actual != test.rlt {
			b.Errorf("Algorithm wrong.")
		}
	}
}
