/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil{
        return false
    }
    
    turtle := head
    hare := head
    
    for hare.Next != nil && hare.Next.Next != nil{
        turtle = turtle.Next
        hare = hare.Next.Next
        if turtle == hare{
            return true
            }
        }
    return false
}