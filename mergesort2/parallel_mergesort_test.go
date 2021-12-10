package mergesort2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayMerge(t *testing.T) {
	tests := []struct {
		unordered []int
		sorted    []int
	}{
		{unordered: []int{15, 8, 13, 30, 21}, sorted: []int{8, 13, 15, 21, 30}},
		{unordered: []int{15, 9}, sorted: []int{9, 15}},
	}
	for k, tt := range tests {

		testname := fmt.Sprintf("sorting test %d", k)
		t.Run(testname, func(t *testing.T) {
			sortedResult := MergeSort(tt.unordered)
			for idx, value := range sortedResult {
				assert.Equal(t, tt.sorted[idx], value)
			}
		})
	}
}
