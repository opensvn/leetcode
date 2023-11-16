/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
type Stack struct {
	elems []int
}

func (s *Stack) Push(v int) {
	s.elems = append(s.elems, v)
}

func (s *Stack) Peek() int {
	return s.elems[len(s.elems)-1]
}

func (s *Stack) Pop() int {
	last := s.Peek()
	s.elems = s.elems[:len(s.elems)-1]
	return last
}

func (s *Stack) Empty() bool {
	return len(s.elems) == 0
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack Stack
	for i := len(nums2)-1; i>=0; i-- {
		v := nums2[i]
		for !stack.Empty() && stack.Peek() <= v {
			stack.Pop()
		}
		if stack.Empty() {
			m[v] = -1
		} else {
			m[v] = stack.Peek()
		}
		stack.Push(v)
	}
	var res []int
	for _, v := range nums1 {
		res = append(res, m[v])
	}
	return res
}
// @lc code=end
