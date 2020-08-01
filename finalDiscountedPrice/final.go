package main

import "fmt"

func FinalDiscountedPrice(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		res[i] = prices[i]
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = prices[index] - prices[i]
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	prices := []int{2, 3, 1, 2, 4, 2}
	fmt.Println(FinalDiscountedPrice(prices))
}