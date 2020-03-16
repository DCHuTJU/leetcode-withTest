package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 一定要有两个指针，一个head指针，一个是移动指针
	var head, rlt *ListNode = nil, nil
	if l1.Val <= l2.Val {
		head = &ListNode{Val: l1.Val}
		l1 = l1.Next
	} else {
		head = &ListNode{Val: l2.Val}
		l2 = l2.Next
	}
	rlt = head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			rlt.Next = &ListNode{
				l2.Val,
				nil,
			}
			l2 = l2.Next
			rlt = rlt.Next
		} else {
			rlt.Next = &ListNode {
				l1.Val,
				nil,
			}
			l1 = l1.Next
			rlt = rlt.Next
		}
	}
	if l1 != nil {
		rlt.Next = l1
	}
	if l2 != nil {
		rlt.Next = l2
	}
	return head
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil &&l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil{
		return l2
	}
	if l1 != nil && l2 == nil{
		return l1
	}
	dummnyhead := &ListNode{Val:-1,Next:nil}
	cur := dummnyhead
	if l1.Val<=l2.Val{
		cur.Next =l1
		cur = cur.Next
		cur.Next = mergeTwoLists2(l1.Next,l2)
	}else{
		cur.Next =l2
		cur = cur.Next
		cur.Next = mergeTwoLists2(l1,l2.Next)
	}
	return dummnyhead.Next
}

func mergeToSlice(list *ListNode) []int {
	rlt := make([]int, 0)
	for list != nil {
		rlt = append(rlt, list.Val)
		list = list.Next
	}
	return rlt
}

func main() {
	var l1 = &ListNode{
		Val:  1,
		Next: &ListNode {
			2,
			&ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	var l2 = &ListNode {
		1,
		&ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	merged := mergeTwoLists1(l1, l2)
	rlt := mergeToSlice(merged)
	fmt.Println(rlt)

	merged = mergeTwoLists2(l1, l2)
	rlt = mergeToSlice(merged)
	fmt.Println(rlt)
}
