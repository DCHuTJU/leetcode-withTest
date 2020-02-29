package main

import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
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

func RecSearchBST(root *TreeNode, want int) int {
	if root == nil {
		return 0
	}
	if root.Val > want {
		return RecSearchBST(root.Left, want)
	} else if root.Val < want {
		return RecSearchBST(root.Right, want)
	} else {
		return root.Val
	}
}

func SearchBST(root *TreeNode, want int) int {
	for root != nil {
		if root.Val == want {
			return root.Val
		} else if root.Val > want {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return 0
}

func main() {
	want := 15
	if SearchBST(Tree, want) != 0 {
		fmt.Println("Find what i want!")
	}
	if RecSearchBST(Tree, want) != 0 {
		fmt.Println("Find what i want!")
	}
}