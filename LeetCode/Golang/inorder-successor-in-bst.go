/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    predecessor := ^int(^uint(0) >> 1)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) == 0 {
            break
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if predecessor == p.Val {
            return root
        }
        predecessor = root.Val
        root = root.Right
    }
    return nil
}
