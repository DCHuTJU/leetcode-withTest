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

var seq_pre = make([]int, 0)

func (root *TreeNode) deleteNode(key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = root.Left.deleteNode(key)
		return root
	}
	if key > root.Val {
		root.Right = root.Right.deleteNode(key)
		return root
	}
	// 找到要删除的目标
	if root.Right == nil {
		//右子树为空
		return root.Left
	}
	if root.Left == nil {
		// 左子树为空
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil {
		// 查找后继，即当 root 的左右子树均不为空时，顺势查找能够替换当前节点的 叶子节点
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right)
	return root
}

func deleteMinNode(root *TreeNode) *TreeNode {
	// 要删除 root 节点
	// 若 root 左子树为空
	if root.Left == nil {
		// 让 root.Right 代替当前 root 的位置
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteMinNode(root.Left)
	return root
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	seq_pre = append(seq_pre, root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func main() {
	rlt := Tree.deleteNode(15)
	// 遍历一下树
	PreOrder(rlt)
	fmt.Println("Result is: ", seq_pre)
}