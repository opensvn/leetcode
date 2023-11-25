/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := make(map[string]int)
	var res []*TreeNode
	var traverse func(*TreeNode) string
	traverse = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		left := traverse(node.Left)
		right := traverse(node.Right)
		s := left + "," + right + "," + strconv.Itoa(node.Val)
		value, ok := memo[s]
		if !ok {
			memo[s] = 1
		} else {
			if value == 1 {
				res = append(res, node)
			}
			memo[s]++
		}
		return s
	}
	traverse(root)
	return res
}
// @lc code=end
