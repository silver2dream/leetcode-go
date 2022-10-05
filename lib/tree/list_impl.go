package tree

import (
	"fmt"
	list "main/lib/linked-list"
)

func NewNode(val any) *Node {
	p := &Node{
		Val: val,
	}
	return p
}

type Node struct {
	Val   any
	Left  *Node
	Right *Node
}

func NewListTree(val []any) ITree {
	p := &ListTree{
		raw: val,
	}
	p.buildTree()
	return p
}

type ListTree struct {
	raw  []any
	root *Node
}

func (p *ListTree) BFSTraversal() {
	traversalQ := list.NewQueue()
	traversalQ.Add(p.root)

	for {
		if traversalQ.Size() <= 0 {
			return
		}

		now := traversalQ.Poll()
		if now.Val.(*Node) == nil {
			continue
		}
		traversalQ.Add(now.Val.(*Node).Left)
		traversalQ.Add(now.Val.(*Node).Right)

		p.print(now.Val.(*Node))
	}
}

func (p *ListTree) MergeSort() {
	//TODO implement me
	panic("implement me")
}

func (p *ListTree) QuickSort() {
	//TODO implement me
	panic("implement me")
}

func (p *ListTree) buildTree() {
	if len(p.raw) < 1 {
		return
	}

	tmp := make([]*Node, 0, cap(p.raw))
	tmp = append(tmp, p.root)
	for i := 1; i <= len(p.raw); i++ {
		now := NewNode(p.raw[i-1])
		tmp = append(tmp, now)
		if i == 1 {
			continue
		}

		parent := i / 2
		if tmp[parent].Left == nil {
			tmp[parent].Left = now
			continue
		}
		if tmp[parent].Right == nil {
			tmp[parent].Right = now
			continue
		}
	}
	p.root = tmp[1]
}

func (p *ListTree) PreorderTraversal() {
	p.preTraversal(p.root)
	fmt.Println()
}

func (p *ListTree) preTraversal(node *Node) {
	if node == nil {
		return
	}

	p.print(node)
	p.preTraversal(node.Left)
	p.preTraversal(node.Right)
}

func (p *ListTree) InorderTraversal() {
	p.inTraversal(p.root)
	fmt.Println()
}

func (p *ListTree) inTraversal(node *Node) {
	if node == nil {
		return
	}

	p.inTraversal(node.Left)
	p.print(node)
	p.inTraversal(node.Right)
}

func (p *ListTree) PostorderTraversal() {
	p.postTraversal(p.root)
	fmt.Println()
}

func (p *ListTree) postTraversal(node *Node) {
	if node == nil {
		return
	}

	p.postTraversal(node.Left)
	p.postTraversal(node.Right)
	p.print(node)
}

func (p *ListTree) print(node *Node) {
	if node == nil || node.Val == nil {
		return
	}
	fmt.Print(node.Val, " ")
}
