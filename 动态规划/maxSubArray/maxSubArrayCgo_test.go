package main

import "testing"

func BenchmarkMaxSubArrayCgo(b *testing.B) {
	tests := struct {
		str  string
		rlt  int
	}{
		"aba",
		0,
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := isHasSubArray(tests.str); actual != tests.rlt {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}