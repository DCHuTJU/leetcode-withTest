package main

import "testing"

func BenchmarkLetterCombinations(b *testing.B) {
	for i:=0; i<b.N; i++ {
		letterCombinations("ab")
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		letterCombinationsv2("ab")
	}
}
