package leetcode

import (
	"fmt"
)

// @leet start
func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)
	for _, row := range grid {
		rowMap[fmt.Sprint(row)]++
	}

	pairs := 0
	for c := range n {
		col := make([]int, n)
		for r := range n {
			col[r] = grid[r][c]
		}

		if count, exists := rowMap[fmt.Sprint(col)]; exists {
			pairs += count
		}
	}

	return pairs
}

// @leet end

