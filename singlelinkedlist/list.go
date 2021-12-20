package slist

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Data int
	Next *ListNode
}
type List struct {
	Head *ListNode
}

func NewEmptyNode() *ListNode {
	return &ListNode{Next: nil}
}
func NewListNode(data int) *ListNode {
	return &ListNode{
		Data: data,
		Next: nil,
	}
}
func NewList() *List {
	return &List{
		Head: NewEmptyNode(),
	}
}
func (l *ListNode) AddTail(data int) {
	for ; l.Next != nil; l = l.Next {
	}
	node := NewListNode(data)
	l.Next = node
}
func (l *ListNode) Search(data int) *ListNode {
	for ; l.Next != nil; l = l.Next {
		if l.Data == data {
			return l
		}
	}
	return nil
}
func (l *ListNode) Delete(node *ListNode) {
	prev := l
	for current := l; current.Next != nil; current = current.Next {
		if current == node {
			prev.Next = current.Next
			current = nil
		}
		prev = current
	}
}
func (l *ListNode) Add(data int) {
	node := NewListNode(data)
	tmp := l.Next
	l.Next = node
	node.Next = tmp
}

/**
* Write a progeam that takes the head of a singly linked list and returns null if there does not exist a cycle
* and the node at the start of the cycle if a cycle is present
 */
// O(N) time, O(1) space
func HasCycleWithNode(head *ListNode) *ListNode {
	history := make(map[*ListNode]bool)
	for current := head; current.Next != nil; current = current.Next {
		if _, ok := history[current.Next]; ok {
			return current
		} else {
			history[current.Next] = true
		}
	}
	return nil
}

// what if we want O(1) space, O(1) times
// the idea is that we have to use two iterators:
// a slow one and a fast one
// In each iteration we advance the slow iterator by one
// and the fast iterator by two.
// the list has cycle if the slow iterator and the fast one will meet.
func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fmt.Printf("slow = %v, fast = %v\n", slow, fast)
		slow, fast = slow.Next, fast.Next.Next
		fmt.Printf("slow = %v, fast = %v\n", slow, fast)
		if slow == fast {
			// there is a cycle
			// we've to find the start of the cycle
			return true
		}
	}
	return false
}

/// Write an algorithm to find the kth to the last element of a single linked list
func FindKToLast(n *ListNode, k int) (int, error) {
	count := 0
	current := n
	for ; current != nil; current = current.Next {
		count++
	}
	pos := count - k
	if pos < 0 {
		return 0, errors.New("k exceeded array size")
	}
	count = 0
	current = n
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	return current.Data, nil
}
