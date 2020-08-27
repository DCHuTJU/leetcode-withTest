package main

import "fmt"

func HeapSort(a []int) []int {
	for i:= len(a)/2 - 1; i>=0; i-- {
		sink(a, i, len(a))
	}

	for i:=len(a)-1; i>=1; i-- {
		swap(a, 0, i)
		sink(a, 0, i)
	}
	return a
}

func sink(a []int, i int, length int) {
	for {
		l := 2 * i + 1
		r := i * 2 + 2
		idx := i
		if l<length && a[l] > a[idx] {
			idx = l
		}
		if r < length && a[r] > a[idx] {
			idx = r
		}
		if idx == i {
			break
		}

		swap(a, i, idx)
		i = idx

	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{1, 5, 2, 1, 3, 4}
	rlt := HeapSort(a)
	fmt.Println(rlt)
}
