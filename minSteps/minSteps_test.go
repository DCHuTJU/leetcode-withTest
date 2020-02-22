package main

import "testing"

func TestMinSteps(t *testing.T) {
	tests := []struct{test, ans int} {
		{
			3,
			3,
		},
		{
			5,
			5,
		},
		{
			30,
			10,
		},
	}
	for _, tt := range tests {
		if actual := minSteps(tt.test); actual != tt.ans {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkMinSteps(b *testing.B) {
	test := struct {
		a int
		b int
	}{
		30,
		10,
	}
	for i:=0; i<b.N; i++ {
		if actual := minSteps(test.a); actual != test.b {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
