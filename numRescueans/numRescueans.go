package main

import (
	"fmt"
	"sort"
)

// 通过查看 cpu.out 可以看出执行过程中 makeslice 将耗费大部分时间，因此可以考虑将 dp 与 grid 重叠进行操作
// 思想与 三角形的那道题一样
func numRescueans(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	i, j := 0, len(people) - 1
	for i<=j {
		if people[i] + people[j] <= limit {
			i++
		}
		j--
		ans++
	}
	return ans
}

func main() {
	var people = []int{1, 2}
	limit := 3
	rlt := numRescueans(people, limit)
	fmt.Println(rlt)
}
