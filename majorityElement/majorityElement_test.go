package main

import "testing"

func BenchmarkMajorityElement1(b *testing.B) {
	var test = struct {
		nums []int
		rlt  int
	}{
		[]int{3,2,3},
		3,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := majorityElement1(test.nums); actual != test.rlt {
			b.Errorf("Benchmark is wrong.")
		}
	}
}

func BenchmarkMajorityElement2(b *testing.B) {
	var test = struct {
		nums []int
		rlt  int
	}{
		[]int{3,2,3},
		3,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := majorityElement2(test.nums); actual != test.rlt {
			b.Errorf("Benchmark is wrong.")
		}
	}
}

func BenchmarkMajorityElement3(b *testing.B) {
	var test = struct {
		nums []int
		rlt  int
	}{
		[]int{3,2,3},
		3,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := majorityElement3(test.nums); actual != test.rlt {
			b.Errorf("Benchmark is wrong.")
		}
	}
}

func BenchmarkMajorityElement4(b *testing.B) {
	var test = struct {
		nums []int
		rlt  int
	}{
		[]int{3,2,3},
		3,
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := majorityElement4(test.nums); actual != test.rlt {
			b.Errorf("Benchmark is wrong.")
		}
	}
}