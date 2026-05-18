package leetcode

import (
	"math"
)

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	longest := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		leftPath := 0
		rightPath := 0

		if node.Left != nil && node.Left.Val == node.Val {
			leftPath = left + 1
		}

		if node.Right != nil && node.Right.Val == node.Val {
			rightPath = right + 1
		}

		if leftPath+rightPath > longest {
			longest = leftPath + rightPath
		}

		return max(leftPath, rightPath)
	}

	dfs(root)

	return longest
}

// @leet end

