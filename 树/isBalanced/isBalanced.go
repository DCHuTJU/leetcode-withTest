package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Tree = &TreeNode {
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if abs(l - r) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Right), maxDepth(root.Left))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	rlt := isBalanced(Tree)
	if rlt == true {
		fmt.Println("Is balanced!")
	} else {
		fmt.Println("Is not balanced!")
	}
}
