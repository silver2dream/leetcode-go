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

func (p *ArrayTree) QuickSort() {
	var sort []any
	for _, val := range p.container {
		if val != nil {
			sort = append(sort, val)
		}
	}
	p.container = sort
	if len(p.container) < 1 {
		return
	}
	p.quickSort(p.container)

	fmt.Println(sort)
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

func (p *ArrayTree) quickSort(sort []any) {
	if len(sort) < 2 {
		return
	}

	length := len(sort)
	pivot := length / 2
	left := 0
	right := length - 2
	sort[pivot], sort[length-1] = sort[length-1], sort[pivot]
	pivot = length - 1
	for {
		for {
			if sort[left].(int) > sort[pivot].(int) || left == right {
				break
			}
			left++
		}

		for {
			if sort[right].(int) < sort[pivot].(int) || left == right {
				break
			}
			right--
		}

		if left == right {
			break
		}
		sort[left], sort[right] = sort[right], sort[left]
	}

	if sort[left].(int) >= sort[pivot].(int) {
		sort[left], sort[pivot] = sort[pivot], sort[left]
		pivot = left
	}

	start := pivot + 1
	end := pivot - 1
	if end < 0 {
		end = 0
	}

	p.quickSort(sort[:end])
	p.quickSort(sort[start:])
}
