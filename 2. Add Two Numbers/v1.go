/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwoNumbersReq(l1, l2, false)
}

func addTwoNumbersReq(l1 *ListNode, l2 *ListNode, isPlusOne bool) *ListNode {
    if l1 == nil && l2 == nil {
        if isPlusOne == true {
            return &ListNode{
                Val: 1,
                Next: nil,
            }
        }

        return nil
    }

    sum := getVal(l1) + getVal(l2)

    if isPlusOne {
        sum = sum + 1
    }

    if sum >= 10 {
        newSum := sum % 10
        
        return &ListNode{
            Val: newSum,
            Next: addTwoNumbersReq(getNext(l1), getNext(l2), true),
        }
    }

    return &ListNode{
        Val: sum,
        Next: addTwoNumbersReq(getNext(l1), getNext(l2), false),
    }
}

func getVal(l *ListNode) int {
    if l == nil {
        return int(0)
    }

    return l.Val
}

func getNext(l *ListNode) *ListNode {
    if l == nil {
        return nil
    }

    return l.Next
}