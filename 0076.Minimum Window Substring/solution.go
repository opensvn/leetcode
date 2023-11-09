/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i, _ := range t {
		need[t[i]]++
	}
	window := make(map[byte]int)
	left, right, valid := 0, 0, 0
	start, minLen := 0, math.MaxInt
	for right < len(s) {
		in := s[right]
		right++
		if _, ok := need[in]; ok {
			window[in]++
			if window[in] == need[in] {
				valid++
			}
		}

		for valid == len(need) {
			if right - left < minLen {
				start = left
				minLen = right - left
			}

			out := s[left]
			left++
			if _, ok := need[out]; ok {
				if window[out] == need[out] {
					valid--
				}
				window[out]--
			}
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[start:start+minLen]
}
// @lc code=end
