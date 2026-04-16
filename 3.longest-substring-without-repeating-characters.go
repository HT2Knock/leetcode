package leetcode

// @leet start
func lengthOfLongestSubstring(s string) int {
	lastSeen := [256]int{}
	start := 0
	maxLength := 0

	for end := 0; end < len(s); end++ {
		char := s[end]

		if start < lastSeen[char] {
			start = lastSeen[char]
		}

		lastSeen[char] = end + 1
		currentLength := end - start + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// @leet end

