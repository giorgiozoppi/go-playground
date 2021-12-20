package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBSTAndCheck(t *testing.T) {
	root := &TreeNode{
		Val:   100,
		Left:  nil,
		Right: nil,
	}
	root.Add(&TreeNode{
		Val:   200,
		Left:  nil,
		Right: nil,
	})
	root.Add(&TreeNode{
		Val:   50,
		Left:  nil,
		Right: nil,
	})
	require.Equal(t, true, IsValidBST(root))
}
