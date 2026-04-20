package leetcode

import (
	"slices"
)

// @leet start
func maxScore(cardPoints []int, k int) int {
	allowCards := slices.Concat(cardPoints[len(cardPoints)-k:], cardPoints[:k])
	maxSum, windowSum := allowCards[0], 0
	start := 0

	for end := 0; end < len(allowCards); end++ {
		windowSum += allowCards[end]

		if end-start+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}

			windowSum -= allowCards[start]
			start++
		}
	}

	return maxSum
}

// @leet end
