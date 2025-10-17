/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    ptr := head
    if ptr == nil {
        return false
    }

    var (
        fastPtr *ListNode
        isExist bool
    )

    fastPtr, isExist = dubbleNext(ptr)
    if !isExist {
        return false
    }

    for {
        if ptr == nil {
            return false
        }

        if ptr == fastPtr {
            return true
        }

        ptr = ptr.Next

        fastPtr, isExist = dubbleNext(fastPtr)
        if !isExist {
            return false
        }
    }

    return false
}

func dubbleNext(ptr *ListNode) (*ListNode, bool) {
    if ptr == nil {
        return nil, false
    }

    if ptr.Next == nil {
        return nil, false
    }

    if ptr.Next.Next == nil {
        return nil, false
    }

    return ptr.Next.Next, true
}