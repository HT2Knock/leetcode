package leetcode

import (
	"container/heap"
	"slices"
)

// @leet start
type ElementHeap struct {
	h      []int
	target int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (e ElementHeap) Len() int {
	return len(e.h)
}

func (e ElementHeap) Less(i, j int) bool {
	diffI := abs(e.h[i] - e.target)
	diffJ := abs(e.h[j] - e.target)

	if diffI != diffJ {
		return diffI > diffJ
	}

	return e.h[i] < e.h[j]
}

func (e ElementHeap) Swap(i, j int) {
	e.h[i], e.h[j] = e.h[j], e.h[i]
}

func (e *ElementHeap) Push(x any) {
	e.h = append(e.h, x.(int))
}

func (e *ElementHeap) Pop() any {
	old := e.h
	n := len(old)
	x := old[n-1]
	e.h = e.h[0 : n-1]

	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	e := &ElementHeap{target: x, h: make([]int, 0, k)}

	for _, ele := range arr {
		if e.Len() < k {
			heap.Push(e, ele)
			continue
		}

		diffEle := abs(ele - e.target)
		diffRoot := abs(e.h[0] - e.target)

		if diffEle < diffRoot || diffEle == diffRoot && ele < e.h[0] {
			e.h[0] = ele
			heap.Fix(e, 0)
		}
	}

	slices.Sort(e.h)

	return e.h
}

// @leet end

