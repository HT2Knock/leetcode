package leetcode

// @leet start
func characterReplacement(s string, k int) int {
	start := 0
	maxLength := 0

	state := [128]int{}
	maxFreq := 0

	for end := 0; end < len(s); end++ {
		state[s[end]]++
		maxFreq = max(maxFreq, state[s[end]])

		for (end-start+1)-maxFreq > k {
			state[s[start]]--
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

// @leet end

