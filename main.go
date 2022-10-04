package main

import (
	"fmt"
	"main/lib/tree"
)

func main() {
	nums := []any{
		5,
		2, 6,
		1, 4, nil, 7,
		nil, nil, 3, nil, nil, nil, nil, nil,
	}
	iTree := tree.NewArrayTree(nums)
	//iTree.PreorderTraversal()
	//iTree.InorderTraversal()
	//iTree.PostorderTraversal()
	//fmt.Println()
	iTreeL := tree.NewListTree(nums)
	//iTreeL.PreorderTraversal()
	//iTreeL.InorderTraversal()
	//iTreeL.PostorderTraversal()
	iTreeL.BFSTraversal()
	fmt.Println()
	iTree.BFSTraversal()
	//iTree.QuickSort()
	//iTree.MergeSort()
}
