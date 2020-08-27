package main

import "fmt"

type KthLargest struct {
	data []int
	maxLength int
	count int
}

func NewNthLargest(k int, nums []int) KthLargest {
	kth := KthLargest{
		data: make([]int, 0),
		maxLength: k,
	}
	for i:=0; i<len(nums); i++ {
		kth.Add(nums[i])
	}
	return kth
}

func (k *KthLargest) Add(num int) int {
	if k.count < k.maxLength - 1 {
		k.data = append(k.data, num)
		k.count += 1
	} else if k.count == k.maxLength - 1 {
		k.data = append(k.data, num)
		k.count += 1
		n:=len(k.data)-1 //data的最大下标
		for i:=n/2; i>=0;i-- {
			heapify(k.data, i)
		}
	} else {
		if num > k.data[0] {
			k.data[0] = num
			heapify(k.data, 0)
		}
	}
	return k.data[0]
}

func heapify(nums []int, idx int) {

	for {
		minPos := idx
		if idx*2 + 1 <= len(nums)-1 && nums[idx*2 + 1] < nums[minPos] {
			minPos = idx * 2 + 1
		}

		if idx*2+2 <= len(nums)-1 && nums[idx*2+2] < nums[minPos] {
			minPos = idx * 2 + 2
		}
		if minPos == idx {
			break
		}
		nums[idx], nums[minPos] = nums[minPos], nums[idx]
		idx = minPos
	}
}

func main() {
	nums := []int{1, 3, 4, 5, 6}
	k := 2
	rlt := NewNthLargest(k, nums)
	fmt.Println(rlt.data)
}
