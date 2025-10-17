/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    var past, current, feature *ListNode


    // init
    past = nil
    current = head
    feature = current.Next


    for {
        current.Next = past
        
        if feature == nil {
            return current
        }

        past = current
        current = feature
        feature = current.Next
    }
}