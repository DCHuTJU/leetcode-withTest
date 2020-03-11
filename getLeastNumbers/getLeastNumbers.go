package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// heap 在 golang 中是封装好的
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{arr[0]}
	}
	if k == len(arr) {
		return arr
	}

	// 构建一个堆
	h := make(IntHeap, 0)
	for i:=0; i<k; i++ {
		h = append(h, arr[i])
	}
	// 对 h 初始化
	heap.Init(&h)
	// fmt.Println(&h)
	for i:=k; i<len(arr); i++ {
		if arr[i] < h[0] {
			h[0] = arr[i]
			sort.Sort(h)
		}
	}
	return h
}

func main() {
	var a = []int{3, 2, 1}
	var k = 2
	rlt := getLeastNumbers(a, k)
	fmt.Println("Result is:", rlt)
}

