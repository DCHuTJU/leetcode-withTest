package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	test := struct {
		tree *TreeNode
		rlt []interface{}
	}{
		&TreeNode{
			3,
			&TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			&TreeNode{
				Val:   20,
				Left:  &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		[]interface{}{3, 9, 20, 15, 7},
	}

	if actual := levelOrder(test.tree); reflect.DeepEqual(actual, test.rlt) == false {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	test := struct {
		tree *TreeNode
		rlt []interface{}
	}{
		&TreeNode{
			3,
			&TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			&TreeNode{
				Val:   20,
				Left:  &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		[]interface{}{3, 9, 20, 15, 7},
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := levelOrder(test.tree); reflect.DeepEqual(actual, test.rlt) == false {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}