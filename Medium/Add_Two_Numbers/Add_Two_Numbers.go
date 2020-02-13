// Runtime: 12 ms, faster than 65.70% of Go online submissions for Add Two Numbers.
// Memory Usage: 5.9 MB, less than 7.32% of Go online submissions for Add Two Numbers.
// 17/12/2019
// DD/MM/YY

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
    "fmt"
    "strconv"
    "math/big"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    if l1 == nil || l1.Val == 0 && l1.Next == nil {
        return l2
    }
    if l2 == nil || l2.Val == 0 && l2.Next == nil {
        return l1
    }
    
    var values1, values2 string
    for l1 != nil {
        values1 += strconv.Itoa(l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        values2 += strconv.Itoa(l2.Val)
        l2 = l2.Next
    }
    
    intValues1 := new(big.Int)
    intValues2 := new(big.Int)
    
    intValues1, _ = intValues1.SetString(Reverse(values1), 10)
    intValues2, _ = intValues2.SetString(Reverse(values2), 10)
    
    sum := new(big.Int)
    sum.Add(intValues1, intValues2)
    final := Reverse(sum.String())
    result := &ListNode{0, nil}
    
    foo, _ := strconv.Atoi(string(final[0]))
    res := &ListNode{Val:foo}
    result.Next = res
    
    for _, i := range final[1:] {
        foo, _ := strconv.Atoi(string(i))
        res.Next = &ListNode{Val:foo}
        res = res.Next
    }
    
    return result.Next
}