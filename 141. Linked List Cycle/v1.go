/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    set := make(map[*ListNode]struct{})

    ptr := head

    for {
        if ptr == nil {
            return false
        } 

        _, ok := set[ptr]
        if ok {
            return true
        }
        
        set[ptr] = struct{}{}

        ptr = ptr.Next
    }

    return false
}