package list

import "fmt"

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
func HasCycle1(head *ListNode) *ListNode {
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

func main() {
	root := NewListNode(3)
	second := NewListNode(5)
	third := NewListNode(8)
	fourth := NewListNode(10)
	root.Next = second
	second.Next = third
	third.Next = root
	fourth.Next = nil
	item := HasCycle(root)
	fmt.Printf("%v\n", item)
}
