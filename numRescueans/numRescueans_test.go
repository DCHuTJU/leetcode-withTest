package main

import "testing"

func TestNumRescueans(t *testing.T) {
	tests := []struct {
		people []int
		limit  int
		rlt    int
	} {
		{
			[]int{1, 2},
			3,
			1,
		},
		{
			[]int{3, 2, 2, 1},
			3,
			3,
		},
		{
			[]int{3, 5, 3, 4},
			5,
			4,
		},
	}

	for _, tt := range tests {
		if actual := numRescueans(tt.people, tt.limit); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkNumRescueans(b *testing.B) {
	test := struct{
		people []int
		limit  int
		rlt    int
	} {
		[]int{3, 5, 3, 4},
		5,
		4,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := numRescueans(test.people, test.limit); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
