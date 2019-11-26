/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var node *TreeNode
    queue := []*TreeNode{root}
    res := make([]int, 0)
    if root == nil {
        return res
    }
    for len(queue) > 0 {
        l := len(queue)
        for i := 0; i < l; i++ {
            node = queue[i]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, node.Val)
        queue = queue[l:]
    }
    return res
}
