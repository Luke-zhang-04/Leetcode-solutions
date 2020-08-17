"""
Runtime: 64 ms, faster than 92.01% of Python3 online submissions for Add Two Numbers.
Memory Usage: 12.8 MB, less than 100.00% of Python3 online submissions for Add Two Numbers.
17/12/2019
DD/MM/YY
"""


class ListNode:
    """
    Definition for singly-linked list.
    Ommit this block in Leetcode
    """

    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:

        if l1 == None or l1.val == 0 and l1.next == None:
            return l2
        if l2 == None or l2.val == 0 and l2.next == None:
            return l1

        values1, values2 = "", ""
        while l1:
            values1 += str(l1.val)
            l1 = l1.next

        while l2:
            values2 += str(l2.val)
            l2 = l2.next

        final = str(int(values1[::-1]) + int(values2[::-1]))[::-1]
        result = ListNode(0)

        res = ListNode(int(final[0]))
        result.next = res

        for i in final[1:]:
            res.next = ListNode(i)
            res = res.next

        return result.next
