/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
type Pair struct {
	value int
	index int
}
func countSmaller(nums []int) []int {
	orig := make([]Pair, len(nums))
	tmp := make([]Pair, len(nums))
	res := make([]int, len(nums))

	for i, num := range nums {
		orig[i] = Pair{num, i}
	}

	var sort func([]Pair, int, int)
	var merge func([]Pair, int, int, int)
	sort = func(pairs []Pair, low, high int) {
		if low >= high {
			return
		}
		mid := low + (high-low)/2
		sort(pairs, low, mid)
		sort(pairs, mid+1, high)
		merge(pairs, low, mid, high)
	}
	merge = func(pairs []Pair, low, mid, high int) {
		for i := low; i <= high; i++ {
			tmp[i] = pairs[i]
		}
		i := low
		j := mid+1
		for k := low; k <= high; k++ {
			if i == mid+1 {
				pairs[k] = tmp[j]
				j++
			} else if j == high+1 {
				pairs[k] = tmp[i]
				i++
				res[pairs[k].index] += j-mid-1
			} else if tmp[i].value > tmp[j].value {
				pairs[k] = tmp[j]
				j++
			} else {
				pairs[k] = tmp[i]
				i++
				res[pairs[k].index] += j-mid-1
			}
		}
	}
	sort(orig, 0, len(orig)-1)
	return res
}
// @lc code=end
