/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) == 0 {
            return 0
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k -= 1
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
    return 0
}
