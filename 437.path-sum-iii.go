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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return sumDfs(root, map[int]int{0: 1}, 0, targetSum)
}

func sumDfs(node *TreeNode, prev map[int]int, sum, target int) int {
	if node == nil {
		return 0
	}

	sum += node.Val

	count := prev[sum-target]

	prev[sum]++

	count += sumDfs(node.Left, prev, sum, target)
	count += sumDfs(node.Right, prev, sum, target)

	prev[sum]--

	return count
}

// @leet end
