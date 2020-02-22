package main

import (
	"reflect"
	"testing"
)

type TestSample struct {
	Matrix [][]int
	Actual []int
}

func TestSpiralOrder(t *testing.T) {
	tests := []TestSample {
		{
			Matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Actual: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			Matrix: [][]int {
				{1, 3, 5},
				{2, 4, 6},
				{3, 6, 9},
			},
			Actual: []int{1, 3, 5, 6, 9, 6, 3, 2, 4},
		},
	}

	for _, tt := range tests {
		if actual := SpiralOrder(tt.Matrix); reflect.DeepEqual(actual, tt.Actual) == false {
			t.Errorf("Algorithm wrong.")
		}
	}
}

/*
F:\algorithm\spiralOrder>go test .
go: cannot find main module; see 'go help modules'

F:\algorithm\spiralOrder>go test SpiralOrder_test.go SpiralOrder.go
ok      command-line-arguments  0.045s
*/

func BenchmarkSpiralOrder(b *testing.B) {
	matrix := [][]int {
		{111, 134, 234},
		{111, 345, 234},
	}
	ans := []int{111, 134, 234, 234, 345, 111}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := SpiralOrder(matrix); reflect.DeepEqual(actual, ans) == false {
			b.Errorf("Benchmark Algorithm wrong.")
		}
		//actual := SpiralOrder(matrix)
		//j := 0
		//for i := 0; i<len(actual) && j<len(ans); i++{
		//	if actual[i] != ans[j] {
		//		b.Errorf("Benchmark Algorithm wrong.")
		//	}
		//	j++
		//}
	}
}

/*
goos: windows
goarch: amd64
BenchmarkSpiralOrder-4   	  934946	      1267 ns/op
PASS
*/

/*
不能正常执行的原因是没有go mod init [packagename]
F:\algorithm\spiralOrder>go test -bench .
goos: windows
goarch: amd64
pkg: algorithm/spiralOrder
BenchmarkSpiralOrder-4           1000000              1042 ns/op
PASS
ok      algorithm/spiralOrder   1.103s
*/
