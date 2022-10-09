package test

import (
	"main/lib/tree"
)

type TreeNode = tree.Node

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return traversalTwoTree(p, q)
}

func traversalTwoTree(p *TreeNode, q *TreeNode) bool {
	isSame := true
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p != nil && q != nil {
		if q.Val != p.Val {
			return false
		}
	}

	isSame = traversalTwoTree(p.Left, q.Left)
	if !isSame {
		return isSame
	}
	isSame = traversalTwoTree(p.Right, q.Right)
	if !isSame {
		return isSame
	}
	return isSame
}
