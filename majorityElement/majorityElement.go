package main

import (
	"fmt"
	"sort"
)

func majorityElement1(nums []int) int {
    var numberMap = make(map[int]int)
    for i:=0; i<len(nums); i++ {
        numberMap[nums[i]]++
    }
    var rlt int
    // type count struct {
    //     val int
    //     num int
    // }
    // var counts = make([]count, 0)
    // for key, value := range numberMap {
    //     tmp := count{}
    //     tmp.val = key
    //     tmp.num = value
    //     counts = append(counts, tmp)
    // }
    // sort.Slice(counts, func(i, j int) bool {
    //     return counts[i].num > counts[j].num
    // })
    // return counts[0].val
    for key, value := range numberMap {
        if value > len(nums) / 2 {
            rlt = key
        }
    }
    return rlt
}

func majorityElement2(nums []int) int {
	var numberMap = make(map[int]int)
	for i:=0; i<len(nums); i++ {
		numberMap[nums[i]]++
	}
	type count struct {
	    val int
	    num int
	}
	var counts = make([]count, 0)
	for key, value := range numberMap {
	    tmp := count{}
	    tmp.val = key
	    tmp.num = value
	    counts = append(counts, tmp)
	}
	sort.Slice(counts, func(i, j int) bool {
	    return counts[i].num > counts[j].num
	})
	return counts[0].val
}

// 摩尔投票法
func majorityElement3(nums []int) int {
	majority := nums[0]
	cnt := 0
	for _, num := range nums {
		if num == majority {
			cnt++
		}else {
			cnt--;
			if cnt==0 {
				majority = num
				cnt = 1
			}
		}
	}
	return majority
}

func majorityElement4(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	var nums = []int{3,2,3}
	fmt.Println("Result is: ", majorityElement1(nums))
	fmt.Println("Result is: ", majorityElement2(nums))
	fmt.Println("Result is: ", majorityElement3(nums))
	fmt.Println("Result is: ", majorityElement4(nums))
}

