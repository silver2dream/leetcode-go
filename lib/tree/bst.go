package tree

import "fmt"

func NewBinarySearchTree(val []any) IBinarySearchTree {
	p := &BinarySearchTree{
		ListTree{raw: val},
	}
	p.buildTree()
	return p
}

type BinarySearchTree struct {
	ListTree
}

func (p *BinarySearchTree) buildTree() {
	for _, val := range p.raw {
		p.Add(val)
	}
}

func (p *BinarySearchTree) Add(val any) {
	afterRoot := p.add(p.root, val)
	if p.root != afterRoot {
		p.root = afterRoot
	}
}

func (p *BinarySearchTree) add(node *Node, val any) *Node {
	if node == nil {
		return NewNode(val)
	}

	if node.Val == val {
		fmt.Println("the same value.")
	} else if node.Val.(int) > val.(int) {
		child := p.add(node.Left, val)
		if node.Left != child {
			node.Left = child
		}
	} else if node.Val.(int) < val.(int) {
		child := p.add(node.Right, val)
		if node.Right != child {
			node.Right = child
		}
	}

	return node
}

func (p *BinarySearchTree) Delete(val any) {
	afterDelete := p.delete(p.root, val)
	if p.root != afterDelete {
		p.root = afterDelete
	}
}

func (p *BinarySearchTree) delete(node *Node, val any) *Node {
	if node == nil {
		return nil
	}

	if node.Val.(int) == val.(int) {
		if node.Right == nil && node.Left == nil {
			return nil
		} else if node.Right != nil && node.Left == nil {
			return node.Right
		} else if node.Left != nil && node.Right == nil {
			return node.Left
		} else if node.Left != nil && node.Right != nil {
			minNode := p.getMin(node.Right)
			minNode.Right = node.Right
			minNode.Left = node.Left
			return minNode
		}
	} else if node.Val.(int) < val.(int) {
		child := p.delete(node.Right, val)
		if node.Right != child {
			node.Right = child
		}
	} else if node.Val.(int) > val.(int) {
		child := p.delete(node.Left, val)
		if node.Left != child {
			node.Left = child
		}
	}

	return node
}

func (p *BinarySearchTree) GetMin() any {
	return p.getMin(p.root)
}

func (p *BinarySearchTree) getMin(root *Node) *Node {
	node := root
	for {
		if node.Left == nil {
			break
		}

		node = node.Left
	}
	return node
}

func (p *BinarySearchTree) GetMax() any {
	node := p.root
	for {
		if node.Right == nil {
			break
		}

		node = node.Right
	}
	return node.Val
}
