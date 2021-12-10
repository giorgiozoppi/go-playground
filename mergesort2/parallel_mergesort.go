/**
** This is the solution to educative.io course in concurrent go.
**/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanIntegers() []int {
	vector := make([]int, 0)
	for {
		fmt.Printf("Digit a number or 'X'>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 {
			fmt.Printf("Invalid data")
			continue
		}
		if text[0] == 'X' {
			break
		}
		v, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("bad number")
			continue
		}
		vector = append(vector, v)
	}
	return vector
}

type outChannel chan []int

func MergeSort(data []int) []int {
	c := make(chan []int)
	PartialSort(data, &c)
	outdata := <-c
	return outdata
}
func Merge(left []int, right []int, output *chan []int) {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			outdata := append(merged, right...)
			*output <- outdata
		} else if len(right) == 0 {
			outdata := append(merged, left...)
			*output <- outdata
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
}
func PartialSort(data []int, output *chan []int) {
	if len(data) <= 1 {
		*output <- data
	}
	mid := len(data) / 2
	go PartialSort(data[:mid], output)
	go PartialSort(data[mid:], output)
	left := <-*output
	right := <-*output
	Merge(left, right, output)
}
