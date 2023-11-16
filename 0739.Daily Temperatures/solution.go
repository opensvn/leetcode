/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
type Element struct {
	Value int
	Index int
}

type Stack struct {
	elems []Element
}

func (s *Stack) Push(v Element) {
	s.elems = append(s.elems, v)
}

func (s *Stack) Peek() Element {
	return s.elems[len(s.elems)-1]
}

func (s *Stack) Pop() Element {
	last := s.Peek()
	s.elems = s.elems[:len(s.elems)-1]
	return last
}

func (s *Stack) Empty() bool {
	return len(s.elems) == 0
}

func dailyTemperatures(temperatures []int) []int {
	var s Stack
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		v := temperatures[i]
		for !s.Empty() && s.Peek().Value <= v {
			s.Pop()
		}
		if s.Empty() {
			res[i] = 0
		} else {
			res[i] = s.Peek().Index - i
		}
		s.Push(Element{Value: v, Index: i})
	}
	return res
}

// @lc code=end
