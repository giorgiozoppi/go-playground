package slist

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

func TestSearchSingleList(t *testing.T) {
	root := NewListNode(20)
	root.AddTail(20)
	root.AddTail(30)
	node := root.Search(20)
	require.NotEqual(t, node, nil)
	require.Equal(t, node.Data, 20)
}
func TestKToLast(t *testing.T) {
	root := NewListNode(20)
	root.AddTail(20)
	root.AddTail(30)
	root.AddTail(50)
	root.AddTail(80)
	found, err := FindKToLast(root, 2)
	require.Equal(t, err, nil)
	require.Equal(t, 50, found)
}
