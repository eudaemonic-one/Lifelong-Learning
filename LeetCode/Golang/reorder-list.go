/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if (head == nil || head.Next == nil || head.Next == nil) {
        return
    }
    // Find the mid point
    var slow = head
    var fast = head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // Reverse the second half in-place
    var prev, cur, next *ListNode
    cur = slow
    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    // Merge in-place
    slow = head
    fast = prev
    for fast != nil && fast.Next != nil {
        slow.Next, slow = fast, slow.Next
        fast.Next, fast = slow, fast.Next
    }
}
