package leetcode

import (
	"cmp"
	"slices"
)

// @leet start
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	merged := make([][]int, 0)

	for _, interval := range intervals {
		if len(merged) > 0 && merged[len(merged)-1][1] >= interval[0] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}

// @leet end

