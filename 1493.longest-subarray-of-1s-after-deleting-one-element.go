package leetcode

// @leet start
func longestSubarray(nums []int) int {
	start := 0
	zeroCount := 0
	maxWindowSize := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}

			start++
		}

		currentWindowSize := (end - start + 1) - 1
		if currentWindowSize > maxWindowSize {
			maxWindowSize = currentWindowSize
		}
	}

	return maxWindowSize
}

// @leet end

