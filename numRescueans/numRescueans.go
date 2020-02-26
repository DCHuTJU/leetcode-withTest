package main

import (
	"fmt"
	"sort"
)

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
