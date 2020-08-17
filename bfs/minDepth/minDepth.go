package main

import (
	"fmt"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

var root = &TreeNode{
	3,
	&TreeNode{
		9,
		nil,
		nil,
	},
	&TreeNode{
		20,
		&TreeNode{
			15,
			nil,
			nil,
		},
		&TreeNode{
			7,
			nil,
			nil,
		},
	},
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// if root.Left == nil || root.Right == nil {
	//     return 1
	// }
	queue := make([]*TreeNode, 0)
	min := 1
	queue = append(queue, root)
	for len(queue) != 0 {
		// 需要层层遍历
		// 每一次存放的都是当前层的内容
		sz := len(queue)
		for i := 0; i < sz; i++ {
			// 判断是否已经达到叶子节点
			if queue[i].Left == nil && queue[i].Right == nil {
				return min
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[sz:]
		min++
	}
	return min
}

func main() {

	rlt := minDepth(root)
	fmt.Println("The result is: ", rlt)
}
