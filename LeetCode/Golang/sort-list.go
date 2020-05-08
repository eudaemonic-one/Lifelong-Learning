/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head
    }
    // Find the mid point
    var slow = head
    var fast = head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // Divide the list into two halfs
    second := slow.Next
    slow.Next = nil
    // Recursively sort two halfs
    l := sortList(head)
    r := sortList(second)
    return merge(l, r)
}

func merge(l, r *ListNode) *ListNode {
    if (l == nil) {
        return r
    } else if (r == nil) {
        return l
    }
    // Let the head be the lesser node between l and r
    if (l.Val > r.Val) {
        l, r = r, l
    }
    head := l
    prev := head
    l = l.Next
    // Merge until either pointer reaches nil pointer
    for l != nil && r != nil {
        if (l.Val < r.Val) {
            prev.Next = l
            l = l.Next
        } else {
            prev.Next = r
            r = r.Next
        }
        prev = prev.Next
    }
    // Concatenate the last nonempty list
    if (l == nil) {
        prev.Next = r
    } else if (r == nil) {
        prev.Next = l
    }
    return head
}
