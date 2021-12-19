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
func TestCreateSingleList(t *testing.T) {
	root := NewListNode(10)
	root.AddTail(20)
	root.AddTail(30)
	require.Equal(t, root.Data, 10)
	require.Equal(t, root.Next.Data, 20)
	require.Equal(t, root.Next.Next.Data, 30)

}
