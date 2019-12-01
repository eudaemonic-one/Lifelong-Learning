/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    node := root
    pVal, qVal := p.Val, q.Val
    for node != nil {
        if pVal < node.Val && qVal < node.Val {
            node = node.Left
        } else if pVal > node.Val && qVal > node.Val {
            node = node.Right
        } else {
            return node
        }
    }
    return nil
}
