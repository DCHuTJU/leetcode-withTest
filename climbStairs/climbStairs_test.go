package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct{
		n 	int
		rlt int
	}{
		{
			3,
			3,
		},
		{
			100,
			1298777728820984005,
		},
	}
	for _, tt := range tests {
		if actual := climbStairs(tt.n); actual != tt.rlt {
			t.Errorf("Algoritm wrong!")
		}
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	var test = struct {
		n int
		rlt int
	}{
		100,
		1298777728820984005,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := climbStairs(test.n); actual != test.rlt {
			b.Errorf("Benchmark sample wrong.")
		}
	}
}