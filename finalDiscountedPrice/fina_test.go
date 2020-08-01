package main

import "testing"

func BenchmarkFinalDiscountedPrice(b *testing.B) {
	prices := []int{2, 3, 1, 2, 4, 2}
	for i:=0; i<b.N; i++ {
		FinalDiscountedPrice(prices)
	}
}
