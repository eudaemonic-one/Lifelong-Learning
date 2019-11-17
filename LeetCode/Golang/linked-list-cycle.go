/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    var slow, fast *ListNode
    if (head == nil) {
        return false
    }
    slow = head
    fast = head.Next
    for slow != fast {
        if (slow.Next == nil || fast.Next == nil || fast.Next.Next == nil) {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}