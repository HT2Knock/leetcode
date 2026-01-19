package leetcode

import (
	"slices"
)

// @leet start
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)

	ops := 0
	start, end := 0, len(nums)-1

	for start < end {
		currentSum := nums[start] + nums[end]

		if currentSum == k {
			ops++
			start++
			end--
		} else if currentSum < k {
			start++
		} else {
			end--
		}
	}

	return ops
}

// @leet end

