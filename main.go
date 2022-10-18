package main

import "fmt"

type node struct {
	black bool
	// Tree fields
	val    uint64
	left   *node
	right  *node
	parent *node
	pp     **node
	// Limit queue fields
	next *node
	prev *node
	// OrderNode
	order *OrderNode
	// This is the other node tying order to another rbtree
	other *node
}

type OrderNode struct {
	priceNode node
	guidNode  node
	quantity  uint64
	stockId   uint64
	kind      int32
	nextFree  *OrderNode
}

type Slab struct {
	free   *OrderNode
	orders []OrderNode
}

func NewSlab(size int) *Slab {
	s := &Slab{orders: make([]OrderNode, size)}
	s.free = &s.orders[0]
	prev := s.free
	for i := 1; i < len(s.orders); i++ {
		curr := &s.orders[i]
		prev.nextFree = curr
		prev = curr
	}
	return s
}

func main() {
	slab := NewSlab(20)
	fmt.Println(slab)
}
