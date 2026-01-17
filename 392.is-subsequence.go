package leetcode

// @leet start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	anchor := 0
	for explorer := 0; explorer < len(t) && anchor < len(s); explorer++ {
		if t[explorer] == s[anchor] {
			anchor++
		}
	}

	return anchor == len(s)
}

// @leet end

