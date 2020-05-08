/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    var slow, fast *ListNode
    if (head == nil) {
        return nil
    }
    slow = head
    fast = head
    for slow != nil && fast != nil {
        slow = slow.Next
        if (slow == nil) {
            return nil
        }
        fast = fast.Next
        if (fast == nil) {
            return nil
        }
        fast = fast.Next
        if (slow == fast) {
            slow = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return slow
        }
    }
    return nil
}
