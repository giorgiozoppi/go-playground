package list

import (
	"fmt"
	"testing"
)

func TestCreateList(t *testing.T) {
	tests := []struct {
		dataset []int
	}{
		{dataset: []int{100, 15, 8, 13, 30, 21, 39, 79, 80, 110}},
		{dataset: []int{200, 35, 52, 31, 53, 28, 54, 70, 80, 110}},
	}
	var root *ListNode = nil
	for k, tt := range tests {
		testcase := fmt.Sprintf("binary search tree test %d\n", k)
		root = nil
		currentData := tt.dataset
		t.Run(testcase, func(t *testing.T) {
			for item := range currentData {
				if root != nil {
					root.AddTail(currentData[item])
				} else {
					root = NewListNode(currentData[item])
				}
			}
		})
	}
}
