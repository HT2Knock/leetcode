package leetcode

// @leet start
func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int, len(arr))
	for _, x := range arr {
		counts[x]++
	}

	seen := make(map[int]struct{}, len(counts))
	for _, freq := range counts {
		if _, exist := seen[freq]; exist {
			return false
		}

		seen[freq] = struct{}{}
	}

	return true
}

// @leet end

