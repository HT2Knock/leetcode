package leetcode

// @leet start
func maximumSubarraySum(nums []int, k int) int64 {
	start, maxSum, windowSum := 0, 0, 0
	seen := make(map[int]bool, len(nums))

	for end := 0; end < len(nums); end++ {
		for seen[nums[end]] || end-start+1 > k {
			windowSum -= nums[start]
			seen[nums[start]] = false
			start++
		}

		seen[nums[end]] = true
		windowSum += nums[end]

		if end-start+1 == k && windowSum > maxSum {
			maxSum = windowSum
		}

	}

	return int64(maxSum)
}

// @leet end
