/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    var dummy = &ListNode{0, head}
    seen := make(map[int]*ListNode)
    prefix := 0
    for cur := dummy; cur != nil; cur = cur.Next {
        prefix += cur.Val
        seen[prefix] = cur
    }
    prefix = 0
    for cur := dummy; cur != nil; cur = cur.Next {
        prefix += cur.Val
        cur.Next = seen[prefix].Next
    }
    return dummy.Next
}
