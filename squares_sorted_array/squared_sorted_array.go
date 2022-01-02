package main

import "sort"

func sortedSquares(nums []int) []int {
	var minor sort.IntSlice
	var major sort.IntSlice
	squared := make([]int, 0, len(nums))
	for _, value := range nums {
		if value < 0 {
			minor = append(minor, value*value)
		} else {
			major = append(major, value*value)
		}
	}
	i := 0
	j := 0
	lastminor := 0
	for i < len(major) && j < len(minor) {
		lastminor = len(minor) - 1 - j
		if major[i] >= minor[lastminor] {
			squared = append(squared, minor[lastminor])
			j++
		} else {
			squared = append(squared, major[i])
			i++
		}
	}
	if i < len(major) {
		for ; i < len(major); i++ {
			squared = append(squared, major[i])
		}
	}
	if j < len(minor) {
		for ; j < len(minor); j++ {
			lastminor = len(minor) - 1 - j
			squared = append(squared, minor[lastminor])
		}
	}
	return squared
}
