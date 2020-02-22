package main

import (
	"reflect"
	"testing"
)

func TestMaxSlideWindow(t *testing.T) {
	// 表格驱动测试
	tests := struct{
		Nums, Rlt []int
		K int
	} {
		[]int{1,3,-1,-3,5,3,6,7},
		[]int{3,3,5,5,6,7},
		3,
	}

	if actual := maxSlidingWindow(tests.Nums, tests.K); reflect.DeepEqual(actual, tests.Rlt) == false {
		t.Errorf("Algorithm wrong.")
	}
}

/*
F:\algorithm\maxSlideWindow> go test .
ok      algorithm/maxSlideWindow        0.051s
*/

func BenchmarkMaxSlideWindow(b *testing.B) {
	nums := []int{1,3,-1,-3,5,3,6,7}
	rlt := []int{3,3,5,5,6,7}
	k := 3

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if ans := maxSlidingWindow(nums, k); reflect.DeepEqual(ans, rlt) == false {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}

/*
F:\algorithm\maxSlideWindow>go test -bench .
goos: windows
goarch: amd64
pkg: algorithm/maxSlideWindow
BenchmarkMaxSlideWindow-4         923686              1205 ns/op
PASS
ok      algorithm/maxSlideWindow        1.184s
*/