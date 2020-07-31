package findMagicIndex

import "testing"

var nums = []int{0, 2, 3, 4, 5}

func TestFindMagicIndex(t *testing.T) {
	if ret := FindMagicIndex(nums); ret != 0 {
		t.Error("Wrong Answer.")
	} else {
		t.Error("Right Answer.")
	}
}

func BenchmarkFindMagicIndex(b *testing.B) {
	for i:=0; i<b.N; i++ {
		FindMagicIndex(nums)
	}
}

func BenchmarkFindMagicIndexII(b *testing.B) {
	for i:=0; i<b.N; i++ {
		FindMagicIndexII(nums)
	}
}