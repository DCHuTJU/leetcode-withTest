package main

import (
	"reflect"
	"testing"
)

func BenchmarkMergeTwoList1(b *testing.B) {
	var test = struct {
		l1  *ListNode
		l2  *ListNode
		rlt []int
	}{
		&ListNode{
			Val:  1,
			Next: &ListNode {
				2,
				&ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		&ListNode {
			1,
			&ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		[]int{1, 1, 2, 3, 3, 4},
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := mergeToSlice(mergeTwoLists1(test.l1, test.l2)); reflect.DeepEqual(actual, test.rlt) == false {
			b.Error("Benchmark wrong. And actual result is:", actual)
		}
	}
}
