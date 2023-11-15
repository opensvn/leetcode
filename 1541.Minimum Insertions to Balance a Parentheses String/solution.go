/*
 * @lc app=leetcode.cn id=1541 lang=golang
 *
 * [1541] 平衡括号字符串的最少插入次数
 */

// @lc code=start
func minInsertions(s string) int {
	var left, right int
	for _, c := range s {
		if c == '(' {
			right += 2
			if right % 2 == 1 {
				left++
				right -= 1
			}
		} else if c == ')' {
			right--
			if right == -1 {
				left++
				right = 1
			}
		}
	}
	return left + right
}
// @lc code=end
