/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    recurse(root, &ans)
    return ans
}

func recurse(root *TreeNode, ans *[]int) {
    if (root == nil) {
        return
    }
    *ans = append(*ans, root.Val)
    recurse(root.Left, ans)
    recurse(root.Right, ans)
}
