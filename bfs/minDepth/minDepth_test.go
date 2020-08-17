package main

import "testing"

func TestMinDepth(t *testing.T) {
	if rlt := minDepth(root); rlt != 2 {
		t.Error("Wrong Answer")
	} else {
		t.Log("Good Job")
	}
}
