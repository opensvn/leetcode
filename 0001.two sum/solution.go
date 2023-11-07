/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	valToIndex := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := valToIndex[diff]; ok {
			return []int{j, i}
		}
		valToIndex[num] = i
	}
	return []int{-1, -1}
}
// @lc code=end
