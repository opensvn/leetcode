/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var result int = math.MinInt
	oneSideMax(root, &result)
	return result
}

func oneSideMax(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	left := maxInt(0, oneSideMax(root.Left, result))
	right := maxInt(0, oneSideMax(root.Right, result))
	*result = maxInt(*result, root.Val+left+right)
	return maxInt(left, right) + root.Val
}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
