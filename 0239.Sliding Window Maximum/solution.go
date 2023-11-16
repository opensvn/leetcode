/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
type Node struct {
	Prev *Node
	Next *Node
	Value int
}

type LinkList struct {
	head *Node
	tail *Node
}

func (l *LinkList) GetFirst() int {
	return l.head.Value
}

func (l *LinkList) PollFirst() {
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head.Next.Prev = nil
		l.head = l.head.Next
	}
}

func (l *LinkList) GetLast() int {
	return l.tail.Value
}

func (l *LinkList) PollLast() {
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail.Prev.Next = nil
		l.tail = l.tail.Prev
	}
}

func (l *LinkList) AddLast(v int) {
	node := &Node{Value: v}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.Prev = l.tail
		l.tail.Next = node
		l.tail = node
	}
}

func (l *LinkList) Empty() bool {
	return l.head == nil && l.tail == nil
}

type MonotonicQueue struct {
	l LinkList
}

func (m *MonotonicQueue) Empty() bool {
	return m.l.Empty()
}

func (m *MonotonicQueue) Push(v int) {
	for (!m.Empty() && m.l.GetLast() < v) {
		m.l.PollLast()
	}
	m.l.AddLast(v)
}

func (m *MonotonicQueue) Pop(v int) {
	if v == m.l.GetFirst() {
		m.l.PollFirst()
	}
}

func (m *MonotonicQueue) Max() int {
	return m.l.GetFirst()
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var m MonotonicQueue
	for i, n := range nums {
		if i < k-1 {
			m.Push(n)
		} else {
			m.Push(n)
			res = append(res, m.Max())
			m.Pop(nums[i-k+1])
		}
	}
	return res
}
// @lc code=end
