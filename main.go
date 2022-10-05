package main

import (
	"fmt"
	"main/lib/tree"
)

func main() {
	//nums := []any{
	//	5,
	//	2, 6,
	//	1, 4, nil, 7,
	//	nil, nil, 3, nil, nil, nil, nil, nil,
	//}
	//iTree := tree.NewArrayTree(nums)
	////iTree.PreorderTraversal()
	////iTree.InorderTraversal()
	////iTree.PostorderTraversal()
	////fmt.Println()
	//iTreeL := tree.NewListTree(nums)
	////iTreeL.PreorderTraversal()
	////iTreeL.InorderTraversal()
	////iTreeL.PostorderTraversal()
	//iTreeL.BFSTraversal()
	//fmt.Println()
	//iTree.BFSTraversal()
	//iTree.QuickSort()
	//iTree.MergeSort()
	//nums := []any{
	//	5,
	//	2, 6,
	//	1, 4, 7, 3,
	//}
	//bst := tree.NewBinarySearchTree(nums)
	//bst.Add(8)
	//fmt.Println(bst.GetMin().(int))
	//fmt.Println(bst.GetMax().(int))
	//bst.SortAsc()
	//fmt.Println()
	//bst.SortDesc()

	num := []any{66, 78, 27, 35, 6, 2, 44, 58, 29, 76}
	bht := tree.NewBinaryHeapTree(num)
	fmt.Println(bht)
	bht.RemoveFromTop()
	fmt.Println(bht)
	bht.AddToBottom(100)
	fmt.Println(bht)
}
