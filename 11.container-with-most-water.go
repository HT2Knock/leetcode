package leetcode

// @leet start
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	largestArea := 0

	for start < end {
		hStart := height[start]
		hEnd := height[end]

		current := (end - start) * min(hEnd, hStart)
		if current > largestArea {
			largestArea = current
		}

		if hStart < hEnd {
			for start < end && height[start] <= hStart {
				start++
			}
		} else {
			for start < end && height[end] <= hEnd {
				end--
			}
		}

	}

	return largestArea
}

// @leet end
