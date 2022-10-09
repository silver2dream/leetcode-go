package tree

type ITree interface {
	GetRoot() *Node
	PreorderTraversal()
	InorderTraversal()
	PostorderTraversal()
	BFSTraversal()
	QuickSort()
	MergeSort()
}

type IBinarySearchTree interface {
	GetRoot() *Node
	Add(val any)
	Delete(val any)
	GetMin() any
	GetMax() any
	SortAsc()
	SortDesc()
}
