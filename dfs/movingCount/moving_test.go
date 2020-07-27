package movingCount

import "testing"

func TestMovingCount(t *testing.T) {
	m, n, k := 2, 3, 1
	if rlt := MovingCount(m, n, k); rlt != 3 {
		t.Fatal("Wrong Answer!")
	}
	t.Log("Right Answer!")
}

func BenchmarkMovingCount(b *testing.B) {
	m, n, k := 2, 3, 1
	for i:=0; i<b.N; i++ {
		MovingCount(m, n, k)
	}
}