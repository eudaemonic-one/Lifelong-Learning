/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    var cnt int
    var dummy = &ListNode{0, head}
    var jump = dummy
    var l = head
    var r = head
    for true {
        // Count the length of the list until k times
        cnt = 0
        for r != nil && cnt < k {
            r = r.Next
            cnt += 1
        }
        if (cnt == k) {
            // Reverse the group of k nodes
            prev, cur := r, l
            for j := 0; j < k; j++ {
                next := cur.Next
                cur.Next = prev
                prev = cur
                cur = next
            }
            // Concatenate the list with the k nodes just reversed 
            jump.Next = prev
            jump = l
            l = r
        } else {
            return dummy.Next
        }
    }
    return nil
}
