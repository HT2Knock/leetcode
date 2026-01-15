package leetcode

import "math"

// @leet start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first := math.MaxInt
	second := math.MaxInt

	for i := range nums {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}

	return false
}

// @leet end
