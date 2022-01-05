package searchsample

import (
	"sort"
)

func SearchSample(value int) bool {
	sampleSlice := []int{1, 6, 19, 34, 21, 82}
	sort.Slice(sampleSlice, func(i, j int) bool {
		return i < j
	})
	pos := sort.SearchInts(sampleSlice, value)
	return sampleSlice[pos] == value
}
