/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	var rank int
	var res int
	traverse(root, k, &rank, &res)
	return res
}

func traverse(root *TreeNode, k int, rank, res *int) {
	if root == nil {
		return
	}

	traverse(root.Left, k, rank, res)
	*rank += 1
	if *rank == k {
		*res = root.Val
		return
	}
	traverse(root.Right, k, rank, res)
}
// @lc code=end
