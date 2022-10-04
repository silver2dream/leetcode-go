package list

type Node struct {
	Val  any
	Next *Node
}

func NewQueue() *Queue {
	p := &Queue{}
	return p
}

type Queue struct {
	head  *Node
	tail  *Node
	count int
}

func (p *Queue) Add(val any) {
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
	p.count++
}

func (p *Queue) Poll() *Node {
	node := p.head
	p.head = p.head.Next
	p.count--
	return node
}

func (p *Queue) Size() int {
	return p.count
}
