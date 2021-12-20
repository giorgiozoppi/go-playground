package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Add(item *TreeNode) {
	var tmp *TreeNode = n
	var parent *TreeNode = tmp
	for tmp != nil {
		if item.Val >= tmp.Val {
			parent = tmp
			tmp = tmp.Right
		} else {
			parent = tmp
			tmp = tmp.Left
		}
	}
	if parent.Val > item.Val {
		parent.Left = item
	} else {
		parent.Right = item
	}
}

func VisitBSTNode(node *TreeNode, visited *[]int) {
	if node == nil {
		return
	}
	VisitBSTNode(node.Left, visited)
	*visited = append(*visited, node.Val)
	VisitBSTNode(node.Right, visited)
}
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	visited := make([]int, 0)
	VisitBSTNode(root, &visited)
	current := visited[0]
	state := true
	for k := 1; k < len(visited) && state; k++ {
		state = state && current < visited[k]
		current = visited[k]

	}
	return state
}
