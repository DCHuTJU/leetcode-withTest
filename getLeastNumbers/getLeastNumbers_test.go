package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	var tests = []struct {
		vals []int
		k    int
		rlt  []int
	}{
		{
			[]int{3, 2, 1},
			2,
			[]int{1, 2},
		},
		{
			[]int{0, 1, 2, 0},
			1,
			[]int{0},
		},
	}
	for _, tt := range tests {
		actual := getLeastNumbers(tt.vals, tt.k)
		// 因为结果可能是倒序也可能是顺序（我们肯定是顺序啦）
		sort.Ints(actual)
		if reflect.DeepEqual(actual, tt.rlt) == false {
			t.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkGetLeastNumbers(b *testing.B) {
	var test = struct {
		vals []int
		k    int
		rlt  []int
	}{
		[]int{3, 2, 1},
		2,
		[]int{1, 2},
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		actual := getLeastNumbers(test.vals, test.k)
		// 因为结果可能是倒序也可能是顺序（我们肯定是顺序啦）
		sort.Ints(actual)
		if reflect.DeepEqual(actual, test.rlt) == false {
			b.Errorf("Benchmar Algorithm wrong.")
		}
	}
}