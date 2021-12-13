package ptree

import (
	"sync"
	"time"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

var wg sync.WaitGroup

func CreateNode(data int, sleep time.Duration) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node) Add(item *Node) {
	var tmp *Node = n
	var parent *Node = tmp
	for tmp != nil {
		if item.Data >= tmp.Data {
			parent = tmp
			tmp = tmp.Right
		} else {
			parent = tmp
			tmp = tmp.Left
		}
	}
	if parent.Data > item.Data {
		parent.Left = item
	} else {
		parent.Right = item
	}
}
func visitTree(root *Node, data *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		visitTree(root.Left, data)
	}
	*data = append(*data, root.Data)
	if root.Right != nil {
		visitTree(root.Right, data)
	}
}
func (n *Node) ProcessNode() []int {
	nodes := make([]int, 0)
	tmp := n
	visitTree(tmp, &nodes)
	return nodes
}
func (n *Node) ProcessNodeParallel(treeTraversal *[]int, mutex *sync.Mutex) {

	defer wg.Done()
	s := NewStack()
	current := n
	if current == nil {
		return
	}
	for current != nil || s.Len() > 0 {
		if current != nil {
			s.Push(current)
			current = current.Left
		} else {
			current = (s.Pop()).(*Node)
			mutex.Lock()
			*treeTraversal = append(*treeTraversal, current.Data)
			mutex.Unlock()
			current = current.Right
		}
	}
}

func (n *Node) TreeTraversalParallel() []int {
	var leftNode *Node
	var dataMutex sync.Mutex
	var rightNode *Node
	leftTree := make([]int, 0)
	rightTree := make([]int, 0)
	leftNode = n.Left
	rightNode = n.Right
	if leftNode != nil {
		wg.Add(1)
		go leftNode.ProcessNodeParallel(&leftTree, &dataMutex)
	}

	if rightNode != nil {
		wg.Add(1)
		go rightNode.ProcessNodeParallel(&rightTree, &dataMutex)
	}
	wg.Wait()
	treeTraversal := make([]int, 0)
	treeTraversal = append(treeTraversal, leftTree...)
	treeTraversal = append(treeTraversal, n.Data)
	treeTraversal = append(treeTraversal, rightTree...)
	return treeTraversal
}
