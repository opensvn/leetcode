/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	var left, right, maxLen int
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := window[c]; ok {
			maxLen = max(maxLen, right - left - 1)
			for s[left] != c {
				delete(window, s[left])
				left++
			}
			left++
		} else {
			window[c]++
		}
	}
	return max(maxLen, right - left)
}
// @lc code=end
