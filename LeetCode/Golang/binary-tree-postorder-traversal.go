/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    recurse(root, &ans)
    return ans
}

func recurse(root *TreeNode, ans *[]int) {
    if (root == nil) {
        return
    }
    recurse(root.Left, ans)
    recurse(root.Right, ans)
    *ans = append(*ans, root.Val)
}
