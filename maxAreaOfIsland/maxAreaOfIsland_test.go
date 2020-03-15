package main

import "testing"

func BenchmarkMaxAreaOfIsland1(b *testing.B) {
	var test = struct {
		area [][]int
		rlt  int
	}{
		[][]int {
			{0,0,1,0,0,0,0,1,0,0,0,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,1,1,0,1,0,0,0,0,0,0,0,0},
			{0,1,0,0,1,1,0,0,1,0,1,0,0},
			{0,1,0,0,1,1,0,0,1,1,1,0,0},
			{0,0,0,0,0,0,0,0,0,0,1,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,0,0,0,0,0,0,1,1,0,0,0,0},
		},
		6,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := maxAreaOfIsland1(test.area); actual != test.rlt {
			b.Error("Benchmark wrong.")
		}
	}
}

func TestMaxAreaOfIsland2(t *testing.T) {
	var test = struct {
		area [][]int
		rlt  int
	}{
		[][]int {
			{0,0,1,0,0,0,0,1,0,0,0,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,1,1,0,1,0,0,0,0,0,0,0,0},
			{0,1,0,0,1,1,0,0,1,0,1,0,0},
			{0,1,0,0,1,1,0,0,1,1,1,0,0},
			{0,0,0,0,0,0,0,0,0,0,1,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,0,0,0,0,0,0,1,1,0,0,0,0},
		},
		6,
	}
	if actual := maxAreaOfIsland2(test.area); actual != test.rlt {
		t.Error("Algorithm is wrong.")
	}
}

func BenchmarkMaxAreaOfIsland2(b *testing.B) {
	var test = struct {
		area [][]int
		rlt  int
	}{
		[][]int {
			{0,0,1,0,0,0,0,1,0,0,0,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,1,1,0,1,0,0,0,0,0,0,0,0},
			{0,1,0,0,1,1,0,0,1,0,1,0,0},
			{0,1,0,0,1,1,0,0,1,1,1,0,0},
			{0,0,0,0,0,0,0,0,0,0,1,0,0},
			{0,0,0,0,0,0,0,1,1,1,0,0,0},
			{0,0,0,0,0,0,0,1,1,0,0,0,0},
		},
		6,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := maxAreaOfIsland2(test.area); actual!=test.rlt {
			b.Errorf("Benchmark wrong, and result is: %d", actual)
		}
	}
}
