package tree

func NewBinaryHeapTree(val []any) *BinaryHeapTree {
	p := &BinaryHeapTree{
		raw: val,
	}
	p.buildTree()
	return p
}

type BinaryHeapTree struct {
	raw      []any
	lastLeaf int
}

func (p *BinaryHeapTree) buildTree() {
	p.lastLeaf = len(p.raw) - 1
	i := p.lastLeaf
	for {
		if i < 0 {
			break
		}

		p.shiftDown(i)
		i--
	}
}

func (p *BinaryHeapTree) shiftDown(i int) {
	if i > p.lastLeaf {
		return
	}

	biggerOne := p.getBiggerOne(i)
	if biggerOne == nil {
		return
	}

	if p.raw[i].(int) < p.raw[biggerOne.(int)].(int) {
		p.raw[i], p.raw[biggerOne.(int)] = p.raw[biggerOne.(int)], p.raw[i]
		p.shiftDown(biggerOne.(int))
	}

}

func (p *BinaryHeapTree) getBiggerOne(i int) any {
	left := i*2 + 1
	right := i*2 + 2

	if left > p.lastLeaf && right > p.lastLeaf {
		return nil
	}

	if left <= p.lastLeaf && right <= p.lastLeaf {
		if p.raw[left].(int) >= p.raw[right].(int) {
			return left
		} else if p.raw[left].(int) < p.raw[right].(int) {
			return right
		}
	} else if left <= p.lastLeaf {
		return left
	} else if right <= p.lastLeaf {
		return right
	}
	return nil
}

func (p *BinaryHeapTree) RemoveFromTop() {
	p.raw[0], p.raw[p.lastLeaf] = p.raw[p.lastLeaf], p.raw[0]
	p.raw = p.raw[:p.lastLeaf]
	p.lastLeaf = p.lastLeaf - 1
	p.shiftDown(0)
}

func (p *BinaryHeapTree) AddToBottom(val any) {
	p.raw = append(p.raw, val)
	p.lastLeaf = len(p.raw) - 1
	i := p.lastLeaf
	p.shiftUp(i)
}

func (p *BinaryHeapTree) shiftUp(i int) {
	if i == 0 || i < 0 {
		return
	}

	parent := (i - 1) / 2
	if p.raw[parent].(int) < p.raw[i].(int) {
		p.raw[parent], p.raw[i] = p.raw[i], p.raw[parent]
		p.shiftUp(parent)
	}
}
