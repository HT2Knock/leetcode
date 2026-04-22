package leetcode

import (
	"cmp"
	"slices"
)

// @leet start
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	end := intervals[0][1]
	nonOverlap := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			nonOverlap++
		}
	}

	return len(intervals) - nonOverlap
}

// @leet end

