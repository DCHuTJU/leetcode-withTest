package main

import (
	"fmt"
	_ "net/http/pprof"
)

var root = &TreeNode{
	1,
	&TreeNode{
		2,
		&TreeNode{
			3,
			nil,
			nil,
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	},
	&TreeNode{
		2,
		&TreeNode{
			4,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	},
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsSymmetric(root.Left, root.Right)
}

func dfsIsSymmetric(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) || (a.Val != b.Val) {
		return false
	}

	return dfsIsSymmetric(a.Left, b.Right) && dfsIsSymmetric(a.Right, b.Left)
}

func main() {
	// http.ListenAndServe("0.0.0.0:9999", nil)
	rlt := IsSymmetric(root)
	fmt.Printf("Result is: %v\n", rlt)
}
