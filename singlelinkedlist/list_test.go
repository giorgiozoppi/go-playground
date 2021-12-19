package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLoopedList(t *testing.T) {
	root := NewListNode(3)
	second := NewListNode(5)
	third := NewListNode(8)
	fourth := NewListNode(10)
	root.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = second
	item := HasCycle(root)
	require.Equal(t, item, true)
}
