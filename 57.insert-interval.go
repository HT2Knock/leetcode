package leetcode

// @leet start
func insert(intervals [][]int, newInterval []int) [][]int {
	merged := make([][]int, 0, len(intervals))
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		merged = append(merged, intervals[i])
		i++
	}

	// Merged the overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}

		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}

		i++
	}

	merged = append(merged, newInterval)
	for j := i; j < n; j++ {
		merged = append(merged, intervals[j])
	}

	return merged
}

// @leet end

