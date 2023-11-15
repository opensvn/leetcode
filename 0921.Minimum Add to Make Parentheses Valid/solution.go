/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */

// @lc code=start
func minAddToMakeValid(s string) int {
	var res, need int
	for _, c := range s {
		if c == '(' {
			need++
		}

		if c == ')' {
			need--
			if need == -1 {
				res++
				need = 0
			}
		}
	}
	return res + need
}
// @lc code=end
