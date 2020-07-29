package main

import (
	"reflect"
	"testing"
)

func TestTreeTraversal(t *testing.T) {
	var test = struct {
		tree *TreeNode
		rlt  int
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
		3,
	}
	if actual := maxDepth(test.tree); actual != test.rlt {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	var test = struct {
		tree *TreeNode
		rlt  int
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
		3,
	}
	for i:=0; i<b.N; i++ {
		if actual := maxDepth(test.tree); actual != test.rlt {
			b.Errorf("Benchamrk algorithm wrong.")
		}
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	var test = struct {
		tree *TreeNode
		rlt  []int
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
		[]int{3, 9, 20, 15, 7},
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		actual := TreeTraversal(test.tree)
		if reflect.DeepEqual(actual, test.rlt) == false {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}

func BenchmarkMaxDepth1(b *testing.B) {
	tree := &TreeNode{
		3,
		&TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		&TreeNode{
			Val: 20,
			Left: &TreeNode{
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
	}
	for i:=0; i<b.N; i++ {
		maxDepth1(tree)
	}
}
