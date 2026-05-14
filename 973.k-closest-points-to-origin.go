package leetcode

import (
	"container/heap"
)

// @leet start
type PointHeap [][]int

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}

	for _, point := range points {
		if h.Len() < k {
			heap.Push(h, point)
			continue
		}

		distP := point[0]*point[0] + point[1]*point[1]
		root := (*h)[0]
		distRoot := root[0]*root[0] + root[1]*root[1]

		if distP < distRoot {
			(*h)[0] = point
			heap.Fix(h, 0)
		}
	}

	return *h
}

// @leet end

