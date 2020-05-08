/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
    res := 0
    helper(root, 0, &res)
    return res
}

func helper(root *TreeNode, parent int, cnt *int) bool {
    if root == nil {
        return true
    }
    l, r := helper(root.Left, root.Val, cnt), helper(root.Right, root.Val, cnt)
    if !(l && r) {
        return false
    }
    *cnt = *cnt + 1
    return root.Val == parent
}
