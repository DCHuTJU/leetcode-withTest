package main

import "testing"

func TestOpenTheLock(t *testing.T) {
	if rlt := openLock(deadends, target); rlt != 1 {
		t.Error("Wrong Answer")
	} else {
		t.Log("Good Job")
	}
}

func BenchmarkOpenTheLock(b *testing.B) {
	for i:=0; i<b.N; i++ {
		openLock(deadends, target)
	}
}
