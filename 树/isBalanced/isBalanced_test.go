package main

import "testing"

var test = struct {
	tree  *TreeNode
	valid bool
}{
	Tree,
	true,
}

func TestIsBalanced(t *testing.T) {

	if actual := isBalanced(test.tree); actual != test.valid {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		if actual := isBalanced(test.tree); actual != test.valid {
			b.Errorf("Benchmark Algorithm wrong.")
		}
	}
}
