package leetcode

// @leet start
func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, a := range piles {
		if a > maxPile {
			maxPile = a
		}
	}

	left, right := 1, maxPile
	for left <= right {
		mid := (left + right) / 2
		tt := 0

		for _, pile := range piles {
			tt += (pile + mid - 1) / mid
		}

		if tt > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

// @leet end

