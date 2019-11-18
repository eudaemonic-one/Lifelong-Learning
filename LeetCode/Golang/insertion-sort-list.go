/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head
    }
    var cur = head
    var dummy = &ListNode{0, nil}
    var prev = dummy
    var next *ListNode
    for cur != nil {
        next = cur.Next
        for prev.Next != nil && prev.Next.Val < cur.Val {
            prev = prev.Next
        }
        cur.Next = prev.Next
        prev.Next = cur
        prev = dummy
        cur = next
    }
    return dummy.Next
}
