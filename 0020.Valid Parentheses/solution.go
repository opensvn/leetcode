/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
type Stack struct {
	elems []string
}

func (s *Stack) Push(str string) {
	s.elems = append(s.elems, str)
}

func (s *Stack) Pop() string {
	if len(s.elems) == 0 {
		return ""
	}
	last := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return last
}

func (s *Stack) Len() int {
	return len(s.elems)
}

func isValid(s string) bool {
	var stack Stack
	for _, e := range s {
		c := string(e)
		if c == "(" || c == "{" || c == "[" {
			stack.Push(c)
		} else if c == ")" {
			last := stack.Pop()
			if last != "(" {
				return false
			}
		} else if c == "}" {
			last := stack.Pop()
			if last != "{" {
				return false
			}
		} else if c == "]" {
			last := stack.Pop()
			if last != "[" {
				return false
			}
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}
// @lc code=end
