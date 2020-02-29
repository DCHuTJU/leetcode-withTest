package main

import "testing"

var test = struct {
	tree *TreeNode
	rlt  bool
}{
	&TreeNode{
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
	},
	true,
}

func TestIsValidBST(t *testing.T) {
	if actual := isValidBST(test.tree); actual != test.rlt {
		t.Errorf("Algorithm is wrong.")
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := isValidBST(test.tree); actual != test.rlt {
			b.Errorf("Algorithm is wrong.")
		}
	}
}
