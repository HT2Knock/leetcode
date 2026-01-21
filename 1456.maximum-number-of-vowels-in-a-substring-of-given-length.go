package leetcode

// @leet start
func maxVowels(s string, k int) int {
	vowels := 0
	isVowelTable := [256]uint8{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}

	for i := 0; i < k; i++ {
		if isVowelTable[s[i]] == 1 {
			vowels++
		}
	}

	maxVowelsCount := vowels
	if maxVowelsCount == k {
		return k
	}

	for i := k; i < len(s); i++ {
		if isVowelTable[s[i]] == 1 {
			vowels++
		}

		if isVowelTable[s[i-k]] == 1 {
			vowels--
		}

		if maxVowelsCount < vowels {
			maxVowelsCount = vowels
			if maxVowelsCount == k {
				return k
			}
		}

	}

	return maxVowelsCount
}

// @leet end
