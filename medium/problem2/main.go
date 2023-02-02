package main

// This implementation uses the existing lists instead of creating a new one

/*
Runtime: 11 ms, beats 72% of Go submissions
Memory Usage: 4.3 MB, beats 100% of Go submissions
2023-02-02
*/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	l1Len := 0
	l2Len := 0
	curNode1 := l1
	curNode2 := l2
	var lastNode1, lastNode2 *ListNode

	for curNode1 != nil || curNode2 != nil {
		total := carry

		if curNode1 != nil {
			total += curNode1.Val
		}
		if curNode2 != nil {
			total += curNode2.Val
		}
		carry = total / 10
		total -= carry * 10

		if curNode1 != nil {
			curNode1.Val = total
			lastNode1 = curNode1
			curNode1 = curNode1.Next
			l1Len++
		}
		if curNode2 != nil {
			curNode2.Val = total
			lastNode2 = curNode2
			curNode2 = curNode2.Next
			l2Len++
		}
	}

	if l1Len > l2Len {
		if carry != 0 {
			newNode := &ListNode{carry, nil}
			lastNode1.Next = newNode
		}

		return l1
	}

	if carry != 0 {
		newNode := &ListNode{carry, nil}
		lastNode2.Next = newNode
	}

	return l2
}
