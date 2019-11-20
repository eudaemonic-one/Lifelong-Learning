/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
    var cnt, length, remainder, width, tmp int
    var cur = root
    var prev, start *ListNode
    ans := make([]*ListNode, 0)
    for cur != nil {
        cur = cur.Next
        cnt += 1
    }
    cur = root
    length = cnt / k
    remainder = cnt % k
    for cur != nil {
        start = cur
        width = length
        if (remainder > 0) {
            width += 1
            remainder -= 1
        }
        for i := 0; i < width; i++ {
            prev = cur
            cur = cur.Next
        }
        prev.Next = nil
        ans = append(ans, start)
        tmp += 1
    }
    for i := tmp; i < k; i++ {
        ans = append(ans, nil)
    }
    return ans
}
