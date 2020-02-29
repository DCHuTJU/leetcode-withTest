package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

// 对树的左右值依次进行对比
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	// 对值进行判断
	if min >= root.Val || max <= root.Val {
		return false
	}
	// 对规则进行判断
	return isBST(root.Left, math.MinInt64, math.MaxInt64) && isBST(root.Right, math.MinInt64, math.MaxInt64)
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
	if isValidBST(tree) == true {
		fmt.Println("This tree is a BST!")
	} else {
		fmt.Println("This tree is not a BST!")
	}
}

