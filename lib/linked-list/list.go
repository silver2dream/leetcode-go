package list

func NewList() *List {
	p := &List{}
	return p
}

type List struct {
	head *Node
	tail *Node
}

func (p *List) GetRoot() *Node {
	return p.head
}

func (p *List) Add(val any) {
	if p.head == nil {
		p.head = &Node{
			Val: val,
		}
		p.tail = p.head
	} else {
		newNode := &Node{
			Val: val,
		}
		p.tail.Next = newNode
		p.tail = newNode
	}
}
