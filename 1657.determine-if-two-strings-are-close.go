package leetcode

import (
	"maps"
	"slices"
)

// @leet start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		m1[word1[i]]++
		m2[word2[i]]++
	}

	for char := range m1 {
		if _, exists := m2[char]; !exists {
			return false
		}
	}

	v1 := slices.Collect(maps.Values(m1))
	v2 := slices.Collect(maps.Values(m2))
	slices.Sort(v1)
	slices.Sort(v2)

	return slices.Equal(v1, v2)
}

// @leet end

