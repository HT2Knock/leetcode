package leetcode

// @leet start
func largestAltitude(gain []int) int {
	current := 0
	highest := 0

	for _, g := range gain {
		current += g
		highest = max(highest, current)
	}

	return highest
}

// @leet end
