package main

import "testing"

var test = struct {
	tree *TreeNode
	want int
}{
	Tree,
	15,
}

func TestSearchBST(t *testing.T) {
	if actual := SearchBST(test.tree, test.want); actual == 0 {
		t.Errorf("Algorithm wrong.")
	}
}

func TestRecSearchBST(t *testing.T) {
	if actual := RecSearchBST(test.tree, test.want); actual == 0 {
		t.Errorf("Algorithm wrong.")
	}
}

func BenchmarkSearchBST(b *testing.B) {
	b.ResetTimer()
	for i:= 0; i<b.N; i++ {
		if actual := SearchBST(test.tree, test.want); actual == 0 {
			b.Errorf("Algorithm wrong.")
		}
	}
}

func BenchmarkRecSearchBST(b *testing.B) {
	b.ResetTimer()
	for i:= 0; i<b.N; i++ {
		if actual := RecSearchBST(test.tree, test.want); actual == 0 {
			b.Errorf("Algorithm wrong.")
		}
	}
}
