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
func goodNodes(root *TreeNode) int {
	return countDfs(root, root.Val)
}

func countDfs(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		maxSoFar = node.Val
		count++
	}

	return countDfs(node.Left, maxSoFar) + countDfs(node.Right, maxSoFar) + count
}

// @leet end

