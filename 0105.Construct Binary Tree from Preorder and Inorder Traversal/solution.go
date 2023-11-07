/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := preorder[0]
	index := 0
	for i, value := range inorder {
		if value == root {
			index = i
			break
		}
	}

	left := inorder[:index]
	right := inorder[index+1:]
	return &TreeNode{
		Val: root,
		Left: buildTree(preorder[1:len(left)+1], left),
		Right: buildTree(preorder[len(left)+1:], right),
	}
}
// @lc code=end
