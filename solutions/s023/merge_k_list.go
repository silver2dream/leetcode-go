package test

import (
	"container/heap"
	list "main/lib/linked-list"
)

type ListNode = list.Node
type mergeHeap []int

func (h mergeHeap) Len() int {
	return len(h)
}

func (h mergeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h mergeHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *mergeHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *mergeHeap) Pop() interface{} {
	old := *h
	tmp := old[len(*h)-1]
	*h = old[0 : len(*h)-1]
	return tmp
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &mergeHeap{}
	heap.Init(h)

	for i := 0; i < len(lists); {
		if lists[i] != nil {
			heap.Push(h, lists[i].Val)
			lists[i] = lists[i].Next
		} else {
			i++
		}
	}

	var head *ListNode
	var tail *ListNode
	for h.Len() > 0 {
		if head == nil {
			head = &ListNode{Val: heap.Pop(h)}
			tail = head
		} else {
			newNode := &ListNode{Val: heap.Pop(h)}
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}
