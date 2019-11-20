/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var p1 = headA
    var p2 = headB
    var l1, l2 int
    for p1 != nil {
        p1 = p1.Next
        l1 += 1
    }
    for p2 != nil {
        p2 = p2.Next
        l2 += 1
    }
    p1, p2 = headA, headB
    for i := l1-l2; i > 0; i-- {
        p1 = p1.Next
    }
    for i := l2-l1; i > 0; i-- {
        p2 = p2.Next
    }
    for p1 != nil && p2 != nil {
        if (p1 == p2) {
            return p1
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    return nil
}
