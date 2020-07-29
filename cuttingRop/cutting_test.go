package cuttingRop

import "testing"

var n = 8

func TestCuttingRop(t *testing.T) {
	if rlt := CuttingRope(n); rlt != 18 {
		t.Error("Wrong Answer!")
	} else {
		t.Log("Success!")
	}
}

func BenchmarkCuttingRop(b *testing.B) {
	for i:=0; i<b.N; i++ {
		CuttingRope(n)
	}

}