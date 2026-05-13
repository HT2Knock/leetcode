package leetcode

// @leet start
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]

	for left <= right {
		mid := (left + right) / 2

		count := 0
		row := n - 1
		col := 0
		for row >= 0 && col < n {
			if matrix[row][col] <= mid {
				count += row + 1
				col++
			} else {
				row--
			}
		}

		if count >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// @leet end

