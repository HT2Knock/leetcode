package leetcode

// @leet start
func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1

	rev := []byte(s)

	for i < j {
		if !isVowel(rev[i]) {
			i++
			continue
		}

		if !isVowel(rev[j]) {
			j--
			continue
		}

		rev[i], rev[j] = rev[j], rev[i]
		i++
		j--
	}

	return string(rev)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// @leet end
