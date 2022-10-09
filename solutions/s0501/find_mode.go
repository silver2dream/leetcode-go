package test

import (
	"main/lib/tree"
)

type TreeNode = tree.Node

var countNow = 0
var countMax = 0
var preVal any
var modes []int

func findMode(root *TreeNode) []int {
	countNow = 0
	countMax = 0
	preVal = nil
	modes = []int{}
	traversal(root)
	return convertArray()
}

func traversal(node *TreeNode) {
	if node == nil {
		return
	}

	traversal(node.Left)
	if preVal == nil {
		preVal = node.Val
		countNow = 1
	} else if preVal != node.Val {
		preVal = node.Val
		countNow = 1
	} else if preVal == node.Val {
		countNow++
	}

	if countNow > countMax {
		countMax = countNow
		modes = []int{}
		modes = append(modes, node.Val.(int))
	} else if countNow == countMax {
		modes = append(modes, node.Val.(int))
	} else if countNow < countMax {

	}

	traversal(node.Right)
}

func convertArray() []int {
	modesAry := make([]int, len(modes))
	for i := 0; i < len(modesAry); i++ {
		modesAry[i] = modes[i]
	}
	return modesAry
}
