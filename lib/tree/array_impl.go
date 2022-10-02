package tree

import "fmt"

type TraversalFunc func(idx int)

func NewArrayTree(val []any) ITree {
	p := &ArrayTree{
		container: val,
	}
	return p
}

// Rules
// root index is zero
// left node index is 2n + 1
// right node index is 2n + 2

type ArrayTree struct {
	container []any
}

func (p *ArrayTree) PreorderTraversal() {
	if len(p.container) < 1 {
		return
	}
	root := 0
	p.preTraversal(root)
	fmt.Println()
}

func (p *ArrayTree) preTraversal(idx int) {
	p.print(p.container[idx])
	p.traversal(idx*2+1, p.preTraversal)
	p.traversal(idx*2+2, p.preTraversal)
}

func (p *ArrayTree) InorderTraversal() {
	if len(p.container) < 1 {
		return
	}
	root := 0
	p.inTraversal(root)
	fmt.Println()
}

func (p *ArrayTree) inTraversal(idx int) {
	p.traversal(idx*2+1, p.inTraversal)
	p.print(p.container[idx])
	p.traversal(idx*2+2, p.inTraversal)
}

func (p *ArrayTree) PostorderTraversal() {
	if len(p.container) < 1 {
		return
	}
	root := 0
	p.postTraversal(root)
	fmt.Println()
}

func (p *ArrayTree) postTraversal(idx int) {
	p.traversal(idx*2+1, p.postTraversal)
	p.traversal(idx*2+2, p.postTraversal)
	p.print(p.container[idx])
}

func (p *ArrayTree) traversal(idx int, funcPtr TraversalFunc) {
	if idx < len(p.container) {
		funcPtr(idx)
	}
	return
}

func (p *ArrayTree) print(val any) {
	if val == nil {
		return
	}
	fmt.Print(val, " ")
}
