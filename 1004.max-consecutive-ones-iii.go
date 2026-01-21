package leetcode

// @leet start
func longestOnes(nums []int, k int) int {
	start := 0
	zerosCount := 0
	maxLen := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zerosCount++
		}

		for zerosCount > k {
			if nums[start] == 0 {
				zerosCount--
			}
			start++
		}

		currentWindowSize := end - start + 1
		if currentWindowSize > maxLen {
			maxLen = currentWindowSize
		}
	}

	return maxLen
}

// @leet end

