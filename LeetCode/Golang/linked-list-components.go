/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
    var ans int
    var cur = head
    set := make(map[int]int)
    for _, val := range G {
        set[val] = 1
    }
    for cur != nil {
        _, ok := set[cur.Val]
        if (ok) {
            cur = cur.Next
            for cur != nil {
                _, ok := set[cur.Val]
                if (!ok) {
                    break
                }
                cur = cur.Next
            }
            ans += 1
        } else {
            cur = cur.Next
        }
    }
    return ans
}
