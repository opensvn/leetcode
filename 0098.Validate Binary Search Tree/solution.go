/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var isValid func(root, min, max *TreeNode) bool
	isValid = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}

		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}

		return isValid(root.Left, min, root) && isValid(root.Right, root, max)
	}
	return isValid(root, nil, nil)
}
// @lc code=end
