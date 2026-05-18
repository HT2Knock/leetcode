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
	paths := make([][]int, 0)

	var dfs func(node *TreeNode, targetSum int, pathNodes []int)
	dfs = func(node *TreeNode, targetSum int, pathNodes []int) {
		if node == nil {
			return
		}

		pathNodes = append(pathNodes, node.Val)
		currentSum := targetSum - node.Val
		if node.Left == nil && node.Right == nil && currentSum == 0 {
			// NOTE: go does slice sharing during recursion
			temp := make([]int, len(pathNodes))
			copy(temp, pathNodes)
			paths = append(paths, temp)
			return
		}

		dfs(node.Left, currentSum, pathNodes)
		dfs(node.Right, currentSum, pathNodes)
	}

	dfs(root, targetSum, []int{})
	return paths
}

// @leet end

