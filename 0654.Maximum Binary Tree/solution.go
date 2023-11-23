/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := -1
	maxIndex := -1
	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}

	leftRoot := constructMaximumBinaryTree(nums[:maxIndex])
	rightRoot := constructMaximumBinaryTree(nums[maxIndex+1:len(nums)])
	return &TreeNode{
		Val: max,
		Left: leftRoot,
		Right: rightRoot,
	}
}
// @lc code=end
