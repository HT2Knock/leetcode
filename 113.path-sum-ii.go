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
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)

	var dfs func(node *TreeNode, targetSum int, path []int)
	dfs = func(node *TreeNode, targetSum int, path []int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		remaining := targetSum - node.Val
		if node.Left == nil && node.Right == nil && remaining == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		dfs(node.Left, remaining, path)
		dfs(node.Right, remaining, path)
	}

	dfs(root, targetSum, []int{})
	return result
}

// @leet end

