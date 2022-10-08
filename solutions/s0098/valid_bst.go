package test

import "main/lib/tree"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type TreeNode = tree.Node

func isValidBST(root *TreeNode) bool {
	var collection []int
	traversal(root, &collection)
	for i := 1; i < len(collection); i++ {
		if collection[i-1] >= collection[i] {
			return false
		}
	}

	return true
}

func traversal(node *TreeNode, collection *[]int) {
	if node == nil {
		return
	}

	traversal(node.Left, collection)
	if node.Val != nil {
		*collection = append(*collection, node.Val.(int))
	}
	traversal(node.Right, collection)
}
