package leetcode

// @leet start
func maxSubArray(nums []int) int {
	start, end := 0, 0
	maxSum := 0

	for end < len(nums) {
		sum := nums[end] - nums[start]
		if sum > maxSum {
			maxSum = end - start + 1
		}
	}

	return maxSum
}

// @leet end
