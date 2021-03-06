/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    curr := res
    carry := 0
    
    for l1 != nil || l2 != nil{
        a, b := 0, 0
        temp := 0
    
    if l1 != nil{
        a = l1.Val
        l1 = l1.Next
    }
    
    if l2 != nil{
        b = l2.Val
        l2 = l2.Next
    }
    
    temp = a + b + carry
    
    if temp >= 10{
        carry = 1
        temp -= 10
    } else{
        carry = 0
    }
    
    curr.Next = &ListNode{Val: temp}
    curr = curr.Next 
    }
    
    if carry == 1{
        curr.Next = &ListNode{Val:1}
    }
    
    return res.Next
}