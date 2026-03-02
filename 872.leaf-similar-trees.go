package leetcode

import (
	"slices"
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1, leaves2 []int

	var collect func(*TreeNode, *[]int)
	collect = func(node *TreeNode, leaves *[]int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			*leaves = append(*leaves, node.Val)
			return
		}

		collect(node.Left, leaves)
		collect(node.Right, leaves)
	}

	collect(root1, &leaves1)
	collect(root2, &leaves2)

	return slices.Equal(leaves1, leaves2)
}

// @leet end

