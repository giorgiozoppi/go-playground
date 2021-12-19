package list

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

// what if we want O(1) space, O(N) times
func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			// there is a cycle
			// we've to find the start of the cycle
			return true
		}
		slow, fast = slow.Next, fast.Next.Next

	}
	return false
}
