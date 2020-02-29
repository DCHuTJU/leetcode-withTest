package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []interface{} {
	var result []interface{}
	if root == nil {
		return result
	}
	// 定义一个双向队列
	queue := list.New()
	queue.PushFront(root)
	// 层次遍历
	for queue.Len() > 0 {
		var current []interface{}
		listlength := queue.Len()
		for i:=0; i<listlength; i++ {
			// 消耗尾部
			// queue.Remove(queue.Back()).(*TreeNode)：移除最后一个元素并将其转化为TreeNode类型
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		for _, v := range current {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	var tree = &TreeNode {
		3,
		&TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		&TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	rlt := levelOrder(tree)
	fmt.Println(rlt)
}