package leetcode

import "math"

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var validBinaryTree func(node *TreeNode, maxVal, minVal int) bool
	validBinaryTree = func(node *TreeNode, maxVal, minVal int) bool {
		if node == nil {
			return true
		}

		if node.Val >= maxVal || node.Val <= minVal {
			return false
		}

		return validBinaryTree(node.Left, node.Val, minVal) && validBinaryTree(node.Right, maxVal, node.Val)
	}

	return validBinaryTree(root, math.MaxInt, math.MinInt)
}

// @leet end
