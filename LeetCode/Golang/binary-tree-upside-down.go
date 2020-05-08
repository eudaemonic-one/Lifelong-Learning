/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    parent := make(map[*TreeNode]*TreeNode)
    node := root
    for node != nil {
        if node.Left != nil {
            parent[node.Left] = node
            node = node.Left
        } else {
            break
        }
    }
    newRoot := node
    for node != nil {
        p, ok := parent[node]
        if ok {
            node.Left = p.Right
            node.Right = p
            node = p
        } else {
            node.Left = nil
            node.Right = nil
            break
        }
    }
    return newRoot
}
