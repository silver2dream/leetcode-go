package test

import "main/lib/tree"

type TreeNode = tree.Node

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderTraversalCollection(root, &result)
	return result
}

func inorderTraversalCollection(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderTraversalCollection(node.Left, result)
	if node.Val != nil {
		*result = append(*result, node.Val.(int))
	}
	inorderTraversalCollection(node.Right, result)
}
