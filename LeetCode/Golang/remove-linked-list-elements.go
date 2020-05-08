/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var dummy = &ListNode{0, head}
    var cur = head
    var prev = dummy
    for cur != nil {
        if (cur.Val != val) {
            prev.Next = cur
            prev = cur
        }
        cur = cur.Next
    }
    prev.Next = cur
    return dummy.Next
}
