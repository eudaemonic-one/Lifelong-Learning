/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    nums := make([]int, 0)
    stack := make([]int, 0)
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    ans := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            ans[stack[len(stack)-1]] = nums[i]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return ans
}
