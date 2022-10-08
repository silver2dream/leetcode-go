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

func (p *ArrayTree) GetRoot() *Node {
	//TODO implement me
	panic("implement me")
}

func (p *ArrayTree) BFSTraversal() {
	//finalLayer := (len(p.container) + 1) / 2
	//for idx, nowLayer := 0, 1; nowLayer < finalLayer; {
	//	if (idx+1)/2 < nowLayer {
	//		p.print(p.container[idx])
	//		idx++
	//	} else {
	//		nowLayer++
	//	}
	//}
	for _, val := range p.container {
		p.print(val)
	}
}

func (p *ArrayTree) MergeSort() {
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
	p.mergeSort(p.container, 0, len(p.container)-1)

	fmt.Println(sort)
}

func (p *ArrayTree) mergeSort(sort []any, start, end int) {
	if start == end {
		return
	}

	pivot := (start + end) / 2
	p.mergeSort(sort, start, pivot)
	p.mergeSort(sort, pivot+1, end)

	var tmp []any
	count := 0
	left := start
	right := pivot + 1

	for count < end-start+1 {
		var leftVal any
		var rightVal any

		if left <= pivot {
			leftVal = sort[left]
		}

		if right <= end {
			rightVal = sort[right]
		}

		if leftVal != nil && rightVal != nil {
			if leftVal.(int) <= rightVal.(int) {
				tmp = append(tmp, leftVal)
				left++
			} else if rightVal.(int) < leftVal.(int) {
				tmp = append(tmp, rightVal)
				right++
			}
		} else if leftVal == nil && rightVal != nil {
			tmp = append(tmp, rightVal)
			right++
		} else if rightVal == nil && leftVal != nil {
			tmp = append(tmp, leftVal)
			left++
		}

		count++
	}

	for i := 0; i < len(tmp); i++ {
		sort[i+start] = tmp[i]
	}

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
