/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := postorder[len(postorder)-1]
	index := 0
	for i, n := range inorder {
		if n == root {
			index = i
			break
		}
	}

	left := inorder[:index]
	right := inorder[index+1:]
	return &TreeNode{
		Val: root,
		Left: buildTree(left, postorder[:len(left)]),
		Right: buildTree(right, postorder[len(left):len(left)+len(right)]),
	}
}
// @lc code=end
