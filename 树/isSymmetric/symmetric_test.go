package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	if IsSymmetric(root) == true {
		t.Log("Right Answer.")
	} else {
		t.Error("Wrong Answer.")
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	for i:=0; i<b.N; i++{
		IsSymmetric(root)
	}
}
