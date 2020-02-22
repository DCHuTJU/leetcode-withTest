package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	queue := []int{}
	result := []int{}
	for i := range nums {
		for i>0 && (len(queue) > 0) && (nums[i] > queue[len(queue) - 1]) {
			// 将比当前元素小的元素祭天
			queue = queue[:len(queue) - 1]
		}
		// 将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			// 维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}

		if i >= k-1 {
			// 放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}

func main() {
	var nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	var k = 3
	rlt := maxSlidingWindow(nums, k)
	fmt.Println(rlt)
}