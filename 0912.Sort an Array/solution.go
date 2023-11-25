/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	tmp := make([]int, len(nums))
	var sort func(nums []int, low, high int)
	var merge func(nums []int, low, mid, high int)
	sort = func(nums []int, low, high int) {
		if low == high {
			return
		}
		mid := low + (high-low)/2
		sort(nums, low, mid)
		sort(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
	merge = func(nums []int, low, mid, high int) {
		for i := low; i <= high; i++ {
			tmp[i] = nums[i]
		}
		i := low
		j := mid+1
		for k := low; k <= high; k++ {
			if i == mid+1 {
				nums[k] = tmp[j]
				j++
			} else if j == high+1 {
				nums[k] = tmp[i]
				i++
			} else if tmp[i] < tmp[j] {
				nums[k] = tmp[i]
				i++
			} else {
				nums[k] = tmp[j]
				j++
			}
		}
	}
	sort(nums, 0, len(nums)-1)
	return nums
}
// @lc code=end
