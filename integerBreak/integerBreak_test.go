package main

import "testing"

func TestIntegerBreak(t *testing.T) {
	tests := []struct{
		val int
		rlt int
	} {
		{
			2,
			1,
		},
		{
			10,
			36,
		},
	}
	for _, tt := range tests {
		if actual := integerBreak(tt.val); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkIntegerBreak(b *testing.B) {
	test := struct{
		val int
		rlt int
	} {
		10,
		36,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := integerBreak(test.val); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}

func BenchmarkIntegerBreakDP(b *testing.B) {
	test := struct{
		val int
		rlt int
	} {
		10,
		36,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := integerBreakDP(test.val); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}