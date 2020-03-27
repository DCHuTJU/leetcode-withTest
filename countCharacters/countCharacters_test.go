package main

import "testing"

func BenchmarkCountCharacters(b *testing.B) {
	var test = struct {
		words []string
		chars string
		rlt   int
	}{
		[]string{"cat", "bt", "hat", "tree"},
		"atach",
		6,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := countCharacters(test.words, test.chars); actual != test.rlt {
			b.Errorf("Benchmark is wrong.\n")
		}
	}
}
