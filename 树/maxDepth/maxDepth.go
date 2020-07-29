package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题主要考虑的是递归和非递归的算法之间的转换

// 递归方法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 非递归实现方法
func TreeTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 用 slice 模拟一个 stack
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	// 为方便检测，把结果换成一个数组输出
	return res
}

// 或者使用层次遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rlt := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		rlt += 1
		for i :=0; i<length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return rlt
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
	rlt := maxDepth(tree)
	fmt.Println(rlt)
	treverse := TreeTraversal(tree)
	fmt.Println(treverse)
	rlt = maxDepth1(tree)
	fmt.Println(rlt)
}