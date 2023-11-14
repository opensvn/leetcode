/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Len int
	getLen(l1, &l1Len)
	var l2Len int
	getLen(l2, &l2Len)

	var shortList *ListNode
	var longList *ListNode

	if l1Len <= l2Len {
		shortList = l1
		longList = l2
	} else {
		shortList = l2
		longList = l1
	}

	var carry int
	res := longList
	pre := longList
	for longList != nil {
		var sum int
		if shortList != nil {
			sum += shortList.Val
			shortList = shortList.Next
		}
		sum += longList.Val + carry
		carry = 0
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		longList.Val = sum
		pre = longList
		longList = longList.Next
	}
	if carry == 1 {
		pre.Next = &ListNode{Val: carry, Next: nil}
	}

	return res
}

func getLen(l *ListNode, count *int) {
	if l == nil || l.Next == nil {
		return
	}

	(*count)++
	getLen(l.Next, count)
}

// @lc code=end
