package binarytree

// TreeNode 定义树节点
type TreeNode struct {
	value       int
	left, right *TreeNode
}

// BinaryTree 定义二叉树
type BinaryTree struct {
	root *TreeNode // 根节点
}


