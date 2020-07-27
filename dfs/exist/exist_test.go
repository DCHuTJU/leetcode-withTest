package exist

import "testing"

var (
	board = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'},
	}
	word = "ABCCED"
)

func TestExist(t *testing.T) {
	if rlt := Exist(board, word); rlt != true {
		t.Fatal("Wrong Answer!")
	}
	t.Log("Right Answer!")
}

func BenchmarkExist(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Exist(board, word)
	}
}
