package ptree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestCreateTree(t *testing.T) {
	tests := []struct {
		dataset []int
	}{
		{dataset: []int{100, 15, 8, 13, 30, 21, 39, 79, 80, 110}},
		{dataset: []int{200, 35, 52, 31, 53, 28, 54, 70, 80, 110}},
	}
	var root *Node = nil

	for k, tt := range tests {
		testcase := fmt.Sprintf("binary search tree test %d\n", k)
		root = nil
		currentData := tt.dataset
		t.Run(testcase, func(t *testing.T) {
			for item := range currentData {
				n := CreateNode(currentData[item], time.Duration(100))
				if root != nil {
					root.Add(n)
				} else {
					root = n
				}
			}
			var expected sort.IntSlice
			expected = currentData[:]
			sort.Stable(expected)
			current := root.ProcessNode()
			for idx, val := range expected {
				assert.Equal(t, val, current[idx])
			}
		})
	}
}
func TestCreateAndParallelTraversal(t *testing.T) {
	tests := []struct {
		dataset []int
	}{
		{dataset: []int{50, 25, 10, 15, 7, 80, 60, 120}},
	}
	var root *Node = nil

	for k, tt := range tests {
		testcase := fmt.Sprintf("binary search tree test %d\n", k)
		root = nil
		currentData := tt.dataset
		t.Run(testcase, func(t *testing.T) {
			for item := range currentData {
				n := CreateNode(currentData[item], time.Duration(100))
				if root != nil {
					root.Add(n)
				} else {
					root = n
				}
			}
			var expected sort.IntSlice
			expected = currentData[:]
			sort.Stable(expected)
			current := root.TreeTraversalParallel()
			for idx, val := range expected {
				assert.Equal(t, val, current[idx])
			}

		})
	}
}
