package tree

type ITree interface {
	PreorderTraversal()
	InorderTraversal()
	PostorderTraversal()
	BFSTraversal()
	QuickSort()
	MergeSort()
}

type IBinarySearchTree interface {
	Add(val any)
	Delete(val any)
	GetMin() any
	GetMax() any
}
