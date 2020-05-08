/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var res *TreeNode
    helper(root, p, q, &res)
    return res
}

func helper(root, p, q *TreeNode, res **TreeNode) bool {
    if root == nil {
        return false
    }
    l := helper(root.Left, p, q, res)
    r := helper(root.Right, p, q, res)
    m := root == p || root == q
    if (l && m) || (m && r) || (l && r) {
        *res = root
    }
    return l || m || r
}
