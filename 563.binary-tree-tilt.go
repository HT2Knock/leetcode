package leetcode

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	tilt := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)

		result := leftDepth - rightDepth
		if result < 0 {
			result = -result
		}
		tilt += result

		return node.Val + leftDepth + rightDepth
	}

	dfs(root)

	return tilt
}

// @leet end

