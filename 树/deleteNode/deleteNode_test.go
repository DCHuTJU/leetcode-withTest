package main

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	test := struct {
		tree *TreeNode
		key  int
		rlt  []int
	}{
		Tree,
		15,
		[]int{3, 9, 20, 7},
	}

	actual := test.tree.deleteNode(test.key)
	PreOrder(actual)
	if reflect.DeepEqual(seq_pre, test.rlt) == false {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkDeleteNode(b *testing.B) {
	test := struct {
		tree *TreeNode
		key  int
		rlt  []int
	}{
		Tree,
		15,
		[]int{3, 9, 20, 7},
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		// 全局数组每一次都要清空
		seq_pre = []int{}
		actual := test.tree.deleteNode(test.key)
		PreOrder(actual)
		if reflect.DeepEqual(seq_pre, test.rlt) == false {
			b.Errorf("Algorithm wrong.")
		}
	}
}