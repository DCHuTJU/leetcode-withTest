package main

import (
	"reflect"
	"testing"
)

func BenchmarkCanEat(b *testing.B) {
	candiesCount := []int{7, 4, 5, 3, 8}
	queries := [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}}
	for i := 0; i < b.N; i++ {
		if rlt := canEat(candiesCount, queries); !reflect.DeepEqual(rlt, []bool{true, false, true}) {
			b.Errorf("Benchmark is wrong.\n")
		}
	}
}
