package main

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	tests := []struct{
		secret string
		guess  string
		rlt    string
	} {
		{
			"1807",
			"7810",
			"1A3B",
		},
		{
			"1123",
			"0111",
			"1A1B",
		},
	}
	for _, tt := range tests {
		if actual := getHint(tt.secret, tt.guess); actual != tt.rlt {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkGetHint(b *testing.B) {
	test := struct {
		secret string
		guess  string
		rlt    string
	}{
		"1807",
		"7810",
		"1A3B",
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := getHint(test.secret, test.guess); actual != test.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
