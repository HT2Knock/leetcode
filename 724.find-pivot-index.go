package leetcode

// @leet start
func pivotIndex(nums []int) int {
	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	leftSum := 0
	for i, x := range nums {
		if leftSum == totalSum-leftSum-x {
			return i
		}

		leftSum += x
	}

	return -1
}

// @leet end

