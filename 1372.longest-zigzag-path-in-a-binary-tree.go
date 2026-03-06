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

func longestZigZag(root *TreeNode) int {
	maxPath := 0

	var dfs func(node *TreeNode, isLeft bool, length int)
	dfs = func(node *TreeNode, isLeft bool, length int) {
		if node == nil {
			return
		}

		if length > maxPath {
			maxPath = length
		}

		if isLeft {
			dfs(node.Right, false, length+1)
			dfs(node.Left, true, 1)
		} else {
			dfs(node.Left, true, length+1)
			dfs(node.Right, false, 1)
		}
	}

	dfs(root.Left, true, 1)
	dfs(root.Right, false, 1)

	return maxPath
}

// @leet end
