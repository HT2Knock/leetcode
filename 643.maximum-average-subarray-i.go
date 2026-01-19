package leetcode

// @leet start
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum

	end := k
	start := end - k

	for end < len(nums) {
		sum += nums[end] - nums[start]
		if sum > maxSum {
			maxSum = sum
		}

		end++
		start++
	}

	return float64(maxSum) / float64(k)
}

// @leet end
